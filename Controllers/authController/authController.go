package authcontroller

import (
	requestDto "CardozoCasariegoLuciano/StudyNotes/Dto/RequestDto"
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	authservice "CardozoCasariegoLuciano/StudyNotes/Service/AuthService"
	customvalidator "CardozoCasariegoLuciano/StudyNotes/helpers/customValidator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authController struct {
	service authservice.IAuthService
}

func NewAuthController() *authController {
	return &authController{service: authservice.NewAuthService()}
}

func (controller authController) Register(c echo.Context) error {
	data := requestDto.RegisterUserDto{}

	if err := c.Bind(&data); err != nil {
		response := responseDto.NewResponse("ERR001", "Error con los datos enviados", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(&data); err != nil {
		response := responseDto.NewResponse("ERR002", "Error en la validacion de los datos enviados", customvalidator.MapValidationErrors(err))
		return c.JSON(http.StatusBadRequest, response)
	}

	//TODO Hacer los tests

	ret := controller.service.RegisterUser(data)
	return c.JSON(http.StatusOK, ret)
}
