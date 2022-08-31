package routes

import (
	users_handler "gozzafadillah/users/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware middleware.JWTConfig
	UserHandler   users_handler.UsersHandler
}

// const server = "https://36e2-2001-448a-1102-1a0f-350a-677f-f95c-668a.ap.ngrok.io/"

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	// access public
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           2592000,
	}))

	e.POST("/login", cl.UserHandler.Login)
	e.POST("/register", cl.UserHandler.Register)

}
