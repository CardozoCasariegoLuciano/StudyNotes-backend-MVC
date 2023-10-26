package usercontroller

import (
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	userservice "CardozoCasariegoLuciano/StudyNotes/Service/UserService"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(c echo.Context) error {
	user := models.User{Name: "Lucho", Password: "123123"}
	return c.JSON(http.StatusOK, user)
}

func All(c echo.Context) error {
	ret := userservice.NewAuthService().ListAll()
	return c.JSON(http.StatusOK, ret)
}
