package main

import (
	users_business "gozzafadillah/users/business"
	con "gozzafadillah/users/config"
	users_handler "gozzafadillah/users/handler"
	"gozzafadillah/users/middlewares"
	users_mysql "gozzafadillah/users/repository/mysql"
	"gozzafadillah/users/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db := con.InitDB()
	con.AutoMigrate(db)

	configJWT := middlewares.ConfigJwt{
		SecretJWT: con.Conf.JWTSecret,
	}

	e := echo.New()

	//Factory
	// Users
	userRepo := users_mysql.NewUsersRepo(db)
	userBusiness := users_business.NewUsersHandler(userRepo, configJWT)
	UserHandler := users_handler.NewUsersHandler(userBusiness)

	routeInit := routes.ControllerList{
		JWTMiddleware: configJWT.Init(),
		UserHandler:   UserHandler,
	}
	routeInit.RouteRegister(e)
	//  start server
	e.Logger.Fatal(e.Start(":8080"))
}
