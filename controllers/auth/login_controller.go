package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"main.go/models/base"
)

func LoginController(c echo.Context) error {
	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "success register",
		Data:    nil,
	})
}
