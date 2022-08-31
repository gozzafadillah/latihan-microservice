package users_domain

import (
	"time"
)

type Users struct {
	UUID      string
	Name      string
	Email     string
	Password  string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	Register(domain Users, file interface{}) error
	Login(email, password string) (string, error)
}

type Repo interface {
	EmailPasswordCheck(email, password string) error
	Store(domain Users) error
	GetUser(email string) (Users, error)
	// Todo: make get user by UUID
}
