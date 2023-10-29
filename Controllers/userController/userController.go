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

// ListAllUsers godoc
// @Summary List all users
// @Description List all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} responseDto.ResponseDto{data=[]responseDto.UserDto}
// @Router /user/all [get]
func (controller *userController) All(c echo.Context) error {
	ret, code := controller.service.ListAll()
	return c.JSON(code, ret)
}
