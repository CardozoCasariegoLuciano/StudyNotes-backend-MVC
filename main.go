package main

import (
	customvalidator "CardozoCasariegoLuciano/StudyNotes/helpers/customValidator"
	"CardozoCasariegoLuciano/StudyNotes/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	//TODO usar los dotenv
	//TODO meter swagger
	//TODO preparar CORS
	port := ":5050"
	e := echo.New()
	e.Validator = customvalidator.NewCustomValidator()

	//Routes
	routes.HanddlerRoutes(e)

	//Starting App
	fmt.Printf("Server runnin on port http://localhost%s", port)
	e.Logger.Fatal(e.Start(port))
}
