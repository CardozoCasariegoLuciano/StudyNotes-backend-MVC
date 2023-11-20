package router

import (
	authcontroller "CardozoCasariegoLuciano/StudyNotes/api/Controllers/authController"
	usercontroller "CardozoCasariegoLuciano/StudyNotes/api/Controllers/userController"
	middlewares "CardozoCasariegoLuciano/StudyNotes/api/Middlewares"

	"github.com/labstack/echo/v4"
)

type Router struct {
	authCtl    authcontroller.AuthController
	userCtl    usercontroller.UserController
	midlewares middlewares.Imiddlewares
}

func NewRouter(
	auth *authcontroller.AuthController,
	user *usercontroller.UserController,
	middle middlewares.Imiddlewares,
) *Router {
	return &Router{
		authCtl:    *auth,
		userCtl:    *user,
		midlewares: middle,
	}
}

func (rt *Router) RegisterRoutes(e *echo.Echo) {
	rt.registerAuthRoutes(e)
	rt.registerUserRoutes(e)
}
