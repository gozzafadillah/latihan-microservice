package main

import (
	product_business "gozzafadillah/products/business"
	con "gozzafadillah/products/config"
	product_handler "gozzafadillah/products/handler"
	"gozzafadillah/products/middlewares"
	product_mysql "gozzafadillah/products/repository/mysql"
	product_routes "gozzafadillah/products/routes"

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
	productRepo := product_mysql.NewProductRepo(db)
	productBusiness := product_business.NewProductBusiness(productRepo)
	productHandler := product_handler.NewProductHandler(productBusiness)

	routeInit := product_routes.ControllerList{
		JWTMiddleware:  configJWT.Init(),
		ProductHandler: productHandler,
	}
	routeInit.RouteRegister(e)
	//  start server
	e.Logger.Fatal(e.Start(":8000"))
}
