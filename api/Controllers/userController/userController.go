package usercontroller

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	userservice "CardozoCasariegoLuciano/StudyNotes/Service/UserService"
	errorcodes "CardozoCasariegoLuciano/StudyNotes/helpers/errorCodes"
	"net/http"
	"strconv"

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
		return c.JSON(http.StatusNotFound, response)
	}

	response := responseDto.NewResponse(
		errorcodes.OK,
		"User loged",
		user,
	)
	return c.JSON(http.StatusOK, response)
}

// Find user godoc
// @Summary Find user by ID
// @Description Find user by ID
// @Tags User
// @Accept json
// @Param userID path string true "ID to find the user"
// @Produce json
// @Success 200 {object} responseDto.ResponseDto{data=responseDto.UserDto}
// @Router /user/{userID} [get]
func (controller *UserController) GetUserByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		response := responseDto.NewResponse(
			errorcodes.INVALID_ID,
			"Invalid ID, must by a number",
			nil,
		)
		return c.JSON(http.StatusNotFound, response)
	}

	user, err := controller.service.FindByID(uint(userID))
	if err != nil {
		response := responseDto.NewResponse(
			errorcodes.NOT_FOUND,
			err.Error(),
			nil,
		)
		return c.JSON(http.StatusNotFound, response)
	}

	response := responseDto.NewResponse(
		errorcodes.OK,
		"User Finded",
		user,
	)
	return c.JSON(http.StatusOK, response)
}
