package product_handler

import (
	product_domain "gozzafadillah/products/domain"
	"gozzafadillah/products/helper/claudinary"
	"gozzafadillah/products/middlewares"
	"net/http"

	product_request "gozzafadillah/products/handler/request"
	helper_users "gozzafadillah/products/helper/users"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductBusiness product_domain.Business
	Validation      *validator.Validate
}

func NewProductHandler(productBusiness product_domain.Business) ProductHandler {
	return ProductHandler{
		ProductBusiness: productBusiness,
		Validation:      validator.New(),
	}
}

// Implementation Create Category
func (ph *ProductHandler) CreateCategory(ctx echo.Context) error {
	req := product_request.Category{}
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": stringerr,
			"status":  http.StatusBadRequest,
		})
	}
	err := ph.ProductBusiness.CreateCategory(product_request.ToDomainCategory(req))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success created",
		"status":  http.StatusCreated,
	})

}

func (ph *ProductHandler) CreateProduct(ctx echo.Context) error {
	req := product_request.Product{}

	parameter := ctx.Param("categorySlug")

	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": stringerr,
			"status":  http.StatusBadRequest,
		})
	}
	// get file
	req.File = claudinary.GetFile(ctx)
	// create data
	err := ph.ProductBusiness.CreateProduct(parameter, product_request.ToDomainProduct(req), req.File)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success created",
		"status":  http.StatusCreated,
	})
}

func (ph *ProductHandler) CreateDetail(ctx echo.Context) error {
	req := product_request.Detail{}
	parameter := ctx.Param("productSlug")
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": stringerr,
			"status":  http.StatusBadRequest,
		})
	}
	err := ph.ProductBusiness.CreateDetail(parameter, product_request.ToDomainDetail(req))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success created",
		"status":  http.StatusCreated,
	})
}

func (ph *ProductHandler) GetCategory(ctx echo.Context) error {
	parameter := ctx.Param("categorySlug")

	res, err := ph.ProductBusiness.GetCategory(parameter)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data",
		"status":  http.StatusOK,
		"result":  res,
	})

}
func (ph *ProductHandler) GetProduct(ctx echo.Context) error {
	parameter := ctx.Param("productSlug")

	res, err := ph.ProductBusiness.GetProduct(parameter)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data",
		"status":  http.StatusOK,
		"result":  res,
	})

}
func (ph *ProductHandler) GetDetail(ctx echo.Context) error {
	parameter := ctx.Param("detailSlug")

	res, err := ph.ProductBusiness.GetDetail(parameter)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data",
		"status":  http.StatusOK,
		"result":  res,
	})

}

// implementation for filter user role by jwt
func (ph *ProductHandler) UserRole(ctx echo.Context, uuid string) string {
	var role string
	jwt := middlewares.GetRaw(ctx)
	claim := middlewares.GetUser(ctx)
	user := helper_users.GetUserUUID(claim.ID, jwt)
	role = user.Result.Role
	return role
}
