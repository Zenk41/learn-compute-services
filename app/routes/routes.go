package routes

import (
	"learn-compute-services/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMIddleware echo.MiddlewareFunc
	JWTMiddleware    middleware.JWTConfig
	AuthController   users.AuthController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	e.Use(cl.LoggerMIddleware)
	users := e.Group("/users")
	users.POST("/register", cl.AuthController.Register)
	users.POST("/login", cl.AuthController.Login)

	withAuth := e.Group("/users", middleware.JWTWithConfig(cl.JWTMiddleware))
	withAuth.GET("", cl.AuthController.GetAllUsers)
	withAuth.POST("", cl.AuthController.CreateUser)
	withAuth.GET("", cl.AuthController.GetUser)

	withAuth.POST("/logout", cl.AuthController.Logout)
}
