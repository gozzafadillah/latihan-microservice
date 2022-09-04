package routes

import (
	users_handler "gozzafadillah/users/handler"
	"gozzafadillah/users/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware middleware.JWTConfig
	UserHandler   users_handler.UsersHandler
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	// log
	middlewares.LogMiddleware(e)

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
	auth := e.Group("users")
	auth.Use(middleware.JWTWithConfig(cl.JWTMiddleware))
	auth.GET("/:id", cl.UserHandler.GetUser)

}
