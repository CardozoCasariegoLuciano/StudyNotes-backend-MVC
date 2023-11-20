package main

import (
	mysql "CardozoCasariegoLuciano/StudyNotes/Repository/MySql"
	authservice "CardozoCasariegoLuciano/StudyNotes/Service/AuthService"
	userservice "CardozoCasariegoLuciano/StudyNotes/Service/UserService"
	"CardozoCasariegoLuciano/StudyNotes/api"
	authcontroller "CardozoCasariegoLuciano/StudyNotes/api/Controllers/authController"
	usercontroller "CardozoCasariegoLuciano/StudyNotes/api/Controllers/userController"
	middlewares "CardozoCasariegoLuciano/StudyNotes/api/Middlewares"
	router "CardozoCasariegoLuciano/StudyNotes/api/Router"
	"CardozoCasariegoLuciano/StudyNotes/helpers/utils"
	"context"

	_ "CardozoCasariegoLuciano/StudyNotes/docs"

	"go.uber.org/fx"
)

// @title			StudyNotes API docTemplate
// @version		1.0
// @BasePath	/api/v1
func main() {
	app := fx.New(
		fx.Provide(
			//Independents
			utils.NewBcrypy,
			utils.NewToken,
			mysql.NewDataBase,

			//Services
			userservice.NewUserService,
			authservice.NewAuthService,

			//Controllers
			authcontroller.NewAuthController,
			usercontroller.NewUserController,
			middlewares.NewMidldeware,

			//Router
			router.NewRouter,

			//API
			api.NewApi,
		),
		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, api *api.Api) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go api.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
