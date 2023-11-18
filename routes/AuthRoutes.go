package routes

import (
	authcontroller "CardozoCasariegoLuciano/StudyNotes/Controllers/authController"
	repository "CardozoCasariegoLuciano/StudyNotes/Repository"
	authservice "CardozoCasariegoLuciano/StudyNotes/Service/AuthService"
	"CardozoCasariegoLuciano/StudyNotes/helpers/utils"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(group *echo.Group, storage repository.IStorage, encripting utils.Ibcrypt, token utils.Itokens) {
	authservice := authservice.NewAuthService(storage, encripting)
	authController := authcontroller.NewAuthController(authservice, token)

	group.POST("/register", authController.Register)
	group.POST("/login", authController.Login)
	group.POST("/logout", authController.Logout)
}
