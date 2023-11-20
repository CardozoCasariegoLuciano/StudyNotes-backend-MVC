package router

import (
	"CardozoCasariegoLuciano/StudyNotes/helpers/utils"

	"github.com/labstack/echo/v4"
)

func (rt *Router) registerUserRoutes(e *echo.Echo) {
	userRoutes := e.Group(utils.BasePath + "/user")

	userRoutes.GET("/me", rt.userCtl.GetUserLoged, rt.midlewares.ValidateToken)
	userRoutes.GET("/all", rt.userCtl.All, rt.midlewares.ValidateToken)
	userRoutes.GET("/:userID", rt.userCtl.All)
}
