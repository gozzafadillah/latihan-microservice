package users_mysql

import (
	users_domain "gozzafadillah/users/domain"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UUID      string
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToDomain(rec Users) users_domain.Users {
	return users_domain.Users{
		UUID:      rec.UUID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Image:     rec.Image,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
