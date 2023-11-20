package usercontroller

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	userservice "CardozoCasariegoLuciano/StudyNotes/Service/UserService"
	errorcodes "CardozoCasariegoLuciano/StudyNotes/helpers/errorCodes"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service userservice.IUserService
}

func NewUserController(service userservice.IUserService) *UserController {
	return &UserController{service: service}
}

// ListAllUsers godoc
// @Summary List all users
// @Description List all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} responseDto.ResponseDto{data=[]responseDto.UserDto}
// @Router /user/all [get]
func (controller *UserController) All(c echo.Context) error {
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

// Find user loged godoc
// @Summary Find user loged
// @Description Find user loged using the cookie
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} responseDto.ResponseDto{data=responseDto.UserDto}
// @Router /user/me [get]
func (controller *UserController) GetUserLoged(c echo.Context) error {
	userID := c.Get("userID")
	user, err := controller.service.FindByID(userID.(uint))
	if err != nil {
		response := responseDto.NewResponse(
			errorcodes.NOT_FOUND,
			err.Error(),
			nil,
		)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responseDto.NewResponse(
		errorcodes.OK,
		"User loged",
		user,
	)
	return c.JSON(http.StatusOK, response)
}
