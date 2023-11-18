package usercontroller

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	userservice "CardozoCasariegoLuciano/StudyNotes/Service/UserService"
	errorcodes "CardozoCasariegoLuciano/StudyNotes/helpers/errorCodes"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	service userservice.IUserService
}

func NewUserController(service userservice.IUserService) *userController {
	return &userController{service: service}
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
	list, err := controller.service.ListAll()
	if err != nil {
		response := responseDto.NewResponse(
			errorcodes.INTERNAL_ERROR,
			err.Error(),
			nil,
		)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responseDto.NewResponse(
		errorcodes.OK,
		"User created",
		map[string][]responseDto.UserDto{
			"list": list,
		},
	)
	return c.JSON(http.StatusOK, response)
}
