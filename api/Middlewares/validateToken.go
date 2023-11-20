package middlewares

import (
	responseDto "CardozoCasariegoLuciano/StudyNotes/Dto/ResponseDto"
	errorcodes "CardozoCasariegoLuciano/StudyNotes/helpers/errorCodes"
	"CardozoCasariegoLuciano/StudyNotes/helpers/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *Middlewares) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(utils.CookieName)

		if err != nil {
			resp := responseDto.NewResponse(errorcodes.NO_TOKEN, "Dont have a token", nil)
			return c.JSON(http.StatusUnauthorized, resp)
		}

		dataToken, err := m.tokenService.ParseToken(cookie.Value)
		if err != nil {
			response := responseDto.NewResponse(errorcodes.WRONG_TOKEN, "Wrong or invalid token", nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		c.Set("userID", dataToken.Id)
		c.Set("userEmail", dataToken.Email)
		c.Set("userRole", dataToken.Role)

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
