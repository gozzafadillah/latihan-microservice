package valid

import (
	product_handler "gozzafadillah/products/handler"
	"gozzafadillah/products/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RoleValidation(role string, productHandler product_handler.ProductHandler) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewares.GetUser(c)
			userRole := productHandler.UserRole(c, claims.ID)

			if userRole == role {
				return hf(c)
			} else {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "Unauthorized account, please contact customer service",
				})
			}
		}
	}
}
