package routes

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	authcontroller "main.go/controllers/auth"
	usercontroller "main.go/controllers/user"
)

func InitRoutes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger()) // untuk memantau user
	e.Use(echojwt.JWT([]byte("123")))
	e.GET("/users", usercontroller.GetUsersController)
	e.POST("/login", authcontroller.LoginController)

}
