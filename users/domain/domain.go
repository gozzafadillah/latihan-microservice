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
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	Register(domain Users, file interface{}) error
	Login(email, password string) (string, error)
	GetUserUUID(uuid string) (Users, error)
	Edit(domain Users, uuid string, file interface{}) error
}

type Repo interface {
	EmailPasswordCheck(email, password string) error
	Store(domain Users) error
	GetUser(email string) (Users, error)
	GetUserUUID(uuid string) (Users, error)
	UpdateUser(domain Users, uuid string) error
}
