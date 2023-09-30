package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"main.go/configs"
	basemodel "main.go/models/base"
	usermmodel "main.go/models/user"
)

func GetUsersController(c echo.Context) error {
	var users []usermmodel.User
	result := configs.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, basemodel.BaseResponse{
			Status:  false,
			Message: "Failed data from database",
			Data:    users,
		})

	}
	return c.JSON(http.StatusOK, basemodel.BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    users,
	})

}
