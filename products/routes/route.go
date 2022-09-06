package product_routes

import (
	product_handler "gozzafadillah/products/handler"
	"gozzafadillah/products/helper/valid"
	"gozzafadillah/products/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	ProductHandler product_handler.ProductHandler
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
	authProduct := e.Group("product")
	authProduct.Use(middleware.JWTWithConfig(cl.JWTMiddleware), valid.RoleValidation("admin", cl.ProductHandler))

	authProduct.POST("/category", cl.ProductHandler.CreateCategory)
	authProduct.POST("/:categorySlug", cl.ProductHandler.CreateProduct)
	authProduct.POST("/detail/:productSlug", cl.ProductHandler.CreateDetail)

	authProduct.GET("/category/:categorySlug", cl.ProductHandler.GetCategory)
	authProduct.GET("/:productSlug", cl.ProductHandler.GetProduct)
	authProduct.GET("/detail/:detailSlug", cl.ProductHandler.GetDetail)
}
