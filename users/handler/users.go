package users_handler

import (
	"fmt"
	users_domain "gozzafadillah/users/domain"
	users_request "gozzafadillah/users/handler/request"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	UsersBusiness users_domain.Business
	Validation    *validator.Validate
}

func NewUsersHandler(userBusiness users_domain.Business) UsersHandler {
	return UsersHandler{
		UsersBusiness: userBusiness,
		Validation:    validator.New(),
	}
}

func (uh *UsersHandler) Register(ctx echo.Context) error {
	req := users_request.UsersJSON{}
	ctx.Bind(&req)
	fmt.Println("data :", req)
	if err := uh.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	// to domain
	err := uh.UsersBusiness.Register(users_request.ToDomain(req))
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

func (uh *UsersHandler) Login(ctx echo.Context) error {
	req := users_request.UsersLoginJSON{}
	ctx.Bind(&req)
	if err := uh.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	token, err := uh.UsersBusiness.Login(req.Email, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"status":  http.StatusOK,
		"result": map[string]interface{}{
			"token": token,
		},
	})
}
