package routes

import (
	mysql "CardozoCasariegoLuciano/StudyNotes/Repository/MySql"
	"CardozoCasariegoLuciano/StudyNotes/helpers/utils"

	"github.com/labstack/echo/v4"
)

func HanddlerRoutes(e *echo.Echo) {
	storage := mysql.NewDataBase()
	encripting := utils.Bcypt{}
	token := utils.Token{}

	authRoutes := e.Group(utils.BasePath + "/auth")
	AuthRoutes(authRoutes, storage, &encripting, &token)

	userRoutes := e.Group(utils.BasePath + "/user")
	UserRoutes(userRoutes)
}
