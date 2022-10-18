package users

import (
	"learn-compute-services/app/middlewares"
	"learn-compute-services/businesses/users"
	"learn-compute-services/controllers"
	"learn-compute-services/controllers/users/request"
	"learn-compute-services/controllers/users/response"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUseCase users.Usecase
}

func NewUserController(authUC users.Usecase) *AuthController {
	return &AuthController{
		authUseCase: authUC,
	}
}

func (ctrl *AuthController) GetAllUsers(c echo.Context) error {
	usersData := ctrl.authUseCase.GetAllUsers()

	users := []response.User{}

	for _, user := range usersData {
		users = append(users, response.FromDomain(user))
	}
	return controllers.NewResponse(c, http.StatusOK, "succes", "all users", users)
}

func (ctrl *AuthController) CreateUser(c echo.Context) error {
	input := request.User{}
	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}
	user := ctrl.authUseCase.CreateUser(input.ToDomain())

	return controllers.NewResponse(c, http.StatusCreated, "succes", "user created", response.FromDomain(user))
}

func (ctrl *AuthController) Register(c echo.Context) error {
	input := request.User{}
	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}
	user := ctrl.authUseCase.CreateUser(input.ToDomain())

	return controllers.NewResponse(c, http.StatusCreated, "succes", "user created", response.FromDomain(user))
}

func (ctrl *AuthController) Login(c echo.Context) error {
	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	if err := userInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	token := ctrl.authUseCase.Login(userInput.ToDomain())
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid email or password",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (ctrl *AuthController) Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	if isListed := middlewares.CheckToken(user.Raw); !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid token",
		})
	}
	middlewares.Logout(user.Raw)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "logout succes",
	})

}

func (ctrl *AuthController) GetUser(c echo.Context) error {
	var id string = c.Param("id")
	user := ctrl.authUseCase.GetUser(id)

	if user.ID == 0 {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "user not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "not found", response.FromDomain(user))
}
