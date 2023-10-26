package routes

import (
	usercontroller "CardozoCasariegoLuciano/StudyNotes/Controllers/userController"

	"github.com/labstack/echo/v4"
)

func UserRoutes(group *echo.Group) {
	group.GET("/me", usercontroller.Init)
	group.GET("/all", usercontroller.All)
	group.GET("/:userID", usercontroller.Init)
}
