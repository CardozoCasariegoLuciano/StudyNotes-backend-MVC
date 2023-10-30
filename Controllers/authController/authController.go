package authcontroller

import (
	requestDto "CardozoCasariegoLuciano/StudyNotes/Dto/RequestDto"
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	authservice "CardozoCasariegoLuciano/StudyNotes/Service/AuthService"
	customvalidator "CardozoCasariegoLuciano/StudyNotes/helpers/customValidator"
	errorcodes "CardozoCasariegoLuciano/StudyNotes/helpers/errorCodes"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authController struct {
	service authservice.IAuthService
}

func NewAuthController() *authController {
	return &authController{service: authservice.NewAuthService()}
}

// Register godoc
// @Summary Register new user
// @Description Charge new user into the database
// @Tags Auth
// @Accept json
// @Param Register body requestDto.RegisterUserDto true "request body"
// @Produce json
// @Success 200 {object} responseDto.ResponseDto{data=swaggertypes.SwaggerCustomTypes{token=string,user=responseDto.UserDto}}
// @Failure 400 {object} responseDto.ResponseDto{data=object}
// @Router /auth/register [post]
func (controller authController) Register(c echo.Context) error {
	data := requestDto.RegisterUserDto{}

	if err := c.Bind(&data); err != nil {
		response := responseDto.NewResponse(
			errorcodes.BODY_TYPES_ERROR,
			"Error con los datos enviados",
			nil,
		)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(&data); err != nil {
		response := responseDto.NewResponse(
			errorcodes.BODY_VALIDATION_ERROR,
			"Error en la validacion de los datos enviados",
			customvalidator.MapValidationErrors(err),
		)
		return c.JSON(http.StatusBadRequest, response)
	}

	ret, code := controller.service.RegisterUser(data)
	return c.JSON(code, ret)
}
