package router

import (
	"CardozoCasariegoLuciano/StudyNotes/helpers/utils"

	"github.com/labstack/echo/v4"
)

func (rt *Router) registerAuthRoutes(e *echo.Echo) {
	authRoutes := e.Group(utils.BasePath + "/auth")

	authRoutes.POST("/register", rt.authCtl.Register)
	authRoutes.POST("/login", rt.authCtl.Login)
	authRoutes.POST("/logout", rt.authCtl.Logout)
}
