package routes

import (
	usercontroller "CardozoCasariegoLuciano/StudyNotes/Controllers/userController"
	mysql "CardozoCasariegoLuciano/StudyNotes/Repository/MySql"
	userservice "CardozoCasariegoLuciano/StudyNotes/Service/UserService"

	"github.com/labstack/echo/v4"
)

func UserRoutes(group *echo.Group) {
	storage := mysql.NewDataBase()
	userservice := userservice.NewUserService(storage)
	userController := usercontroller.NewUserController(userservice)

	group.GET("/me", userController.All)
	group.GET("/all", userController.All)
	group.GET("/:userID", userController.All)
}
