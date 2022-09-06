package users_request

import users_domain "gozzafadillah/users/domain"

type UsersJSON struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Role     string
	Image    string      `json:"img" form:"img"`
	File     interface{} `json:"file,omitempty"`
}

func ToDomain(req UsersJSON) users_domain.Users {
	return users_domain.Users{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Image:    req.Image,
		Role:     req.Role,
	}
}

type UsersLoginJSON struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

func ToDomainLogin(req UsersLoginJSON) users_domain.Users {
	return users_domain.Users{
		Email:    req.Email,
		Password: req.Password,
	}
}
