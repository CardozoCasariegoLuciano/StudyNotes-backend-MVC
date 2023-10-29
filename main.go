package main

import (
	"CardozoCasariegoLuciano/StudyNotes/configuration"
	customvalidator "CardozoCasariegoLuciano/StudyNotes/helpers/customValidator"
	"CardozoCasariegoLuciano/StudyNotes/routes"
	"fmt"

	_ "CardozoCasariegoLuciano/StudyNotes/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			StudyNotes API docTemplate
// @version		1.0
// @BasePath	/api/v1
func main() {
	//TODO prepara el repo con los Actions y todo eso
	//TODO meter mysql

	//Get config
	config := configuration.GetConfig()
	port := config.App.Port

	//Init echo
	e := echo.New()
	e.Validator = customvalidator.NewCustomValidator()

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

	//Middleware
	e.Use(middleware.Recover())
	if config.App.Logger {
		e.Use(middleware.Logger())
	}

	//Swager
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//Routes
	routes.HanddlerRoutes(e)

	//Starting App
	fmt.Printf("Server runnin on port http://localhost%s", port)
	e.Logger.Fatal(e.Start(port))
}
