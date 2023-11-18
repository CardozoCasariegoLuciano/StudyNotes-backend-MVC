package authcontroller

import (
	requestDto "CardozoCasariegoLuciano/StudyNotes/Dto/RequestDto"
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	authservice "CardozoCasariegoLuciano/StudyNotes/Service/AuthService"
	customvalidator "CardozoCasariegoLuciano/StudyNotes/helpers/customValidator"
	errorcodes "CardozoCasariegoLuciano/StudyNotes/helpers/errorCodes"
	errortypes "CardozoCasariegoLuciano/StudyNotes/helpers/errorTypes"
	"CardozoCasariegoLuciano/StudyNotes/helpers/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	service authservice.IAuthService
	token   utils.Itokens
}

func NewAuthController(service authservice.IAuthService, token utils.Itokens) *AuthController {
	return &AuthController{service: service, token: token}
}

// Register godoc
// @Summary Register new user
// @Description Charge new user into the database and set a cookie with de JWT
// @Tags Auth
// @Accept json
// @Param Register body requestDto.RegisterUserDto true "request body"
// @Produce json
// @Success 200 {object} responseDto.ResponseDto{data=swaggertypes.SwaggerCustomTypes{user=responseDto.UserDto}}
// @Failure 400 {object} responseDto.ResponseDto{data=object}
// @Router /auth/register [post]
func (controller *AuthController) Register(c echo.Context) error {
	data := requestDto.RegisterUserDto{}

	if err := utils.BindBody(c, &data); err != nil {
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

	createdUser, err := controller.service.RegisterUser(data)
	if err != nil {
		if err == errortypes.MailAlreadyTaken {
			response := responseDto.NewResponse(
				errorcodes.MAIL_TAKEN,
				err.Error(),
				nil,
			)
			return c.JSON(http.StatusBadRequest, response)
		}

		response := responseDto.NewResponse(
			errorcodes.INTERNAL_ERROR,
			err.Error(),
			nil,
		)
		return c.JSON(http.StatusInternalServerError, response)
	}

	//Generate Token
	token, err := controller.token.GenerateToken(*createdUser)
	if err != nil {
		response := responseDto.NewResponse(errorcodes.JWT_ERROR, "trouble creating a JWT", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	//cookie := http.Cookie{
	//Name:     utils.CookieName,
	//Value:    token,
	//Path:     utils.BasePath,
	//Secure:   true,
	//SameSite: http.SameSiteNoneMode,
	//HttpOnly: true,
	//}
	c.SetCookie(utils.CreateCookie(token))

	response := responseDto.NewResponse(
		errorcodes.OK,
		"User created",
		map[string]*responseDto.UserDto{
			"user": createdUser,
		},
	)
	return c.JSON(http.StatusCreated, response)
}

// Login godoc
// @Summary Login user
// @Description Login user and set the cookie
// @Tags Auth
// @Accept json
// @Param Login body requestDto.LoginUserDto true "request body"
// @Produce json
// @Success 200 {object} responseDto.ResponseDto{data=swaggertypes.SwaggerCustomTypes{user=responseDto.UserDto}}
// @Failure 400 {object} responseDto.ResponseDto{data=object}
// @Router /auth/login [post]
func (controller *AuthController) Login(c echo.Context) error {
	data := requestDto.LoginUserDto{}

	if err := utils.BindBody(c, &data); err != nil {
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

	userLogued, err := controller.service.LoginUser(data)
	if err != nil {
		if err == errortypes.WrongPassOrEmail {
			response := responseDto.NewResponse(
				errorcodes.WRONG_LOGIN_DATA,
				err.Error(),
				nil,
			)
			return c.JSON(http.StatusBadRequest, response)
		}

		response := responseDto.NewResponse(
			errorcodes.INTERNAL_ERROR,
			err.Error(),
			nil,
		)
		return c.JSON(http.StatusInternalServerError, response)
	}

	//Generate Token
	t, err := controller.token.GenerateToken(*userLogued)
	if err != nil {
		response := responseDto.NewResponse(errorcodes.JWT_ERROR, "trouble creating a JWT", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	c.SetCookie(utils.CreateCookie(t))

	response := responseDto.NewResponse(
		errorcodes.OK,
		"User loged",
		map[string]*responseDto.UserDto{
			"user": userLogued,
		},
	)
	return c.JSON(http.StatusOK, response)
}

// Logout godoc
// @Summary Logout user
// @Description Remove cookie
// @Tags Auth
// @Success 200
// @Router /auth/logout [post]
func (controller *AuthController) Logout(c echo.Context) error {
	_, err := c.Cookie(utils.CookieName)
	if err != nil {
		return nil
	}

	c.SetCookie(utils.DeleteCookie())
	return nil
}
