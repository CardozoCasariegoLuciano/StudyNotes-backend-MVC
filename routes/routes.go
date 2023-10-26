package routes

import "github.com/labstack/echo/v4"

func HanddlerRoutes(e *echo.Echo) {
	basePath := "/api/v1"

	authRoutes := e.Group(basePath + "/auth")
	AuthRoutes(authRoutes)

	userRoutes := e.Group(basePath + "/user")
	UserRoutes(userRoutes)
}
