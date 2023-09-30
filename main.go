package main

import (
	"github.com/labstack/echo/v4"
	"main.go/configs"
	"main.go/routes"
)

func main() {
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoutes(e)
	e.Start(":8000")
}
