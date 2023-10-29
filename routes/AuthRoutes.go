package routes

import (
	authcontroller "CardozoCasariegoLuciano/StudyNotes/Controllers/authController"
	usercontroller "CardozoCasariegoLuciano/StudyNotes/Controllers/userController"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(group *echo.Group) {
	authController := authcontroller.NewAuthController()

	group.POST("/register", authController.Register)
	group.POST("/login", usercontroller.Init)
}
