package main

import (
	"encoding/json"
	_userUseCase "learn-compute-services/businesses/users"
	_userController "learn-compute-services/controllers/users"
	"learn-compute-services/controllers/users/request"
	_driverFactory "learn-compute-services/drivers"
	"testing"

	"net/http"

	_middleware "learn-compute-services/app/middlewares"
	_routes "learn-compute-services/app/routes"
	_dbDriver "learn-compute-services/drivers/mysql"
	"learn-compute-services/drivers/mysql/users"
	"learn-compute-services/util"

	echo "github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
)

func mainTest() *echo.Echo {

	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}
	db := configDB.InitDB()
	_dbDriver.DBMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:      util.GetConfig("JWT_SECRET_KEY"),
		ExpireDuration: 24,
	}

	configLogger := _middleware.ConfigLogger{
		Format: "[${time_rfc3339}] method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
	}

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo, &configJWT)
	userCtrl := _userController.NewUserController(userUseCase)

	routesInit := _routes.ControllerList{
		LoggerMIddleware: configLogger.Init(),
		JWTMiddleware:    configJWT.Init(),
		AuthController:   *userCtrl,
	}

	routesInit.RouteRegister(e)
	return e
}

// Cleaning up seeds on database
func cleanUp(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		configDB := _dbDriver.ConfigDB{
			DB_USERNAME: util.GetConfig("DB_USERNAME"),
			DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
			DB_HOST:     util.GetConfig("DB_HOST"),
			DB_PORT:     util.GetConfig("DB_PORT"),
			DB_NAME:     util.GetConfig("DB_TEST_NAME"),
		}
		db := configDB.InitDB()
		_dbDriver.CleanSeeds(db)
	}
}

func getJWTToken(t *testing.T) string {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}
	db := configDB.InitDB()
	user := _dbDriver.SeedUser(db)

	var userRequest *request.User = &request.User{
		Email:    user.Email,
		Password: user.Password,
	}

	var res *http.Response = apitest.New().
		Handler(mainTest()).
		Post("/users").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End().Response

	var response map[string]string = map[string]string{}

	json.NewDecoder(res.Body).Decode(&response)

	var token string = response["token"]
	var JWT_TOKEN = "Bearer " + token

	return JWT_TOKEN
}

func getUser() users.User {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()
	user := _dbDriver.SeedUser(db)

	return user
}

func TestRegister_Success(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "testing123@testing.com",
		Password: "testing123",
	}
	apitest.New().
		Observe(cleanUp).
		Handler(mainTest()).
		Post("/users/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()

}

func TestRegister_ValidationFailed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "",
		Password: "",
	}
	apitest.New().
		Handler(mainTest()).
		Post("/users/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()

}

func TestLogin_Success(t *testing.T) {
	user := getUser()
	var userRequest *request.User = &request.User{
		Email:    user.Email,
		Password: user.Password,
	}
	apitest.New().
		Handler(mainTest()).
		Post("/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()

}

func TestLogin_ValidationFailed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "",
		Password: "",
	}
	apitest.New().
		Handler(mainTest()).
		Post("/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

}

func TestLogin_Failed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "doesntexist@gmail.com",
		Password: "itsfailed",
	}
	apitest.New().
		Handler(mainTest()).
		Post("/users/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusUnauthorized).
		End()

}

func TestCreateUser_Success(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "testcreate@gmai.com",
		Password: "testcreate",
	}

	var token string = getJWTToken(t)
	apitest.New().
		Observe(cleanUp).
		Handler(mainTest()).
		Post("users").
		Header("Authorization", token).
		JSON(userRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()

}

func TestGetAllUsers_Success(t *testing.T) {
	var token string = getJWTToken(t)
	apitest.New().
		Observe(cleanUp).
		Handler(mainTest()).
		Get("users").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}
