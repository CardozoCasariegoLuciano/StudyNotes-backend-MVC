package usercontroller

import (
	models "CardozoCasariegoLuciano/StudyNotes/Models"
	userservice "CardozoCasariegoLuciano/StudyNotes/Service/UserService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	service userservice.IUserService
}

func NewUserController() *userController {
	return &userController{service: userservice.NewUserService()}
}

func Init(c echo.Context) error {
	user := models.User{Name: "Lucho", Password: "123123"}
	return c.JSON(http.StatusOK, user)
}

func (controller *userController) All(c echo.Context) error {
	ret := controller.service.ListAll()
	return c.JSON(http.StatusOK, ret)
}
