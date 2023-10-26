package authcontroller

import (
	requestDto "CardozoCasariegoLuciano/StudyNotes/Dto/RequestDto"
	authservice "CardozoCasariegoLuciano/StudyNotes/Service/AuthService"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	data := requestDto.RegisterUserDto{}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, "mal")
	}
	//TODO Valdiar la data SEGUIR POR ACA

	//TODO responder bien segun corresponda
	//TODO Hacer los tests

	ret := authservice.NewAuthService().RegisterUser(data)
	return c.JSON(http.StatusOK, ret)
}
