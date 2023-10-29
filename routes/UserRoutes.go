package routes

import (
	usercontroller "CardozoCasariegoLuciano/StudyNotes/Controllers/userController"

	"github.com/labstack/echo/v4"
)

func UserRoutes(group *echo.Group) {
	userController := usercontroller.NewUserController()

	group.GET("/me", usercontroller.Init)
	group.GET("/all", userController.All)
	group.GET("/:userID", usercontroller.Init)
}
