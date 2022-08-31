package users_business

import (
	"errors"
	users_domain "gozzafadillah/users/domain"
	"gozzafadillah/users/middlewares"

	"github.com/google/uuid"
)

type UsersBusiness struct {
	JWT       middlewares.ConfigJwt
	UsersRepo users_domain.Repo
}

// Login implements users_domain.Business
func (ub UsersBusiness) Login(email string, password string) (string, error) {
	// Check email and Password
	err := ub.UsersRepo.EmailPasswordCheck(email, password)
	if err != nil {
		return "", errors.New("email and password miss match")
	}
	// Get User by email
	usersData, err := ub.UsersRepo.GetUser(email)
	if err != nil {
		return "", errors.New("user not found")
	}
	// generate token JWT
	token, err := ub.JWT.GenerateToken(usersData.UUID, usersData.Email)
	if err != nil {
		return "", errors.New("failed generate token")
	}
	return token, nil
}

// Register implements users_domain.Business
func (ub UsersBusiness) Register(domain users_domain.Users) error {
	// make uuid
	uuidData := uuid.New()
	domain.UUID = uuidData.String()
	// store data
	err := ub.UsersRepo.Store(domain)
	if err != nil {
		return errors.New("failed store data")
	}
	return nil
}

func NewUsersHandler(userRepo users_domain.Repo, jwt middlewares.ConfigJwt) users_domain.Business {
	return UsersBusiness{
		JWT:       jwt,
		UsersRepo: userRepo,
	}
}
