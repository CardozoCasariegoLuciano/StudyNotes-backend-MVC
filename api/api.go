package api

import (
	router "CardozoCasariegoLuciano/StudyNotes/api/Router"
	"CardozoCasariegoLuciano/StudyNotes/configuration"
	customvalidator "CardozoCasariegoLuciano/StudyNotes/helpers/customValidator"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Api struct {
	router router.Router
}

func NewApi(router *router.Router) *Api {
	return &Api{
		router: *router,
	}
}

func (api *Api) Start() {
	//Get config
	config := configuration.GetConfig()
	port := config.App.Port

	//Init echo
	e := echo.New()
	e.Validator = customvalidator.NewCustomValidator()

	//CORS
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowCredentials: true,
		},
	))

	//Middleware
	e.Use(middleware.Recover())
	if config.App.Logger {
		e.Use(middleware.Logger())
	}

	//Swager
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//Routes
	api.router.RegisterRoutes(e)

	//Starting App
	fmt.Printf("Server runnin on port http://localhost%s", port)
	e.Logger.Fatal(e.Start(port))
}
