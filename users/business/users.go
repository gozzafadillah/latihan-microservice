package users_business

import (
	"errors"
	"fmt"
	users_domain "gozzafadillah/users/domain"
	"gozzafadillah/users/helper/claudinary"
	"gozzafadillah/users/middlewares"

	"github.com/google/uuid"
)

type UsersBusiness struct {
	JWT       middlewares.ConfigJwt
	UsersRepo users_domain.Repo
}

// GetUserUUID implements users_domain.Business
func (ub UsersBusiness) GetUserUUID(uuid string) (users_domain.Users, error) {
	dataUser, err := ub.UsersRepo.GetUserUUID(uuid)
	if err != nil {
		return users_domain.Users{}, err
	}
	return dataUser, nil
}

// Edit implements users_domain.Business
func (ub UsersBusiness) Edit(domain users_domain.Users, uuid string, file interface{}) error {

	// upload image
	img, _ := claudinary.ImageUploadHelper(file, "users")

	domain.Image = img
	if domain.Image == "" {
		domain.Image = "https://res.cloudinary.com/dt91kxctr/image/upload/v1655825545/go-bayeue/users/download_o1yrxx.png"
	}
	// update data
	err := ub.UsersRepo.UpdateUser(domain, uuid)
	if err != nil {
		return err
	}
	return nil
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
	token, err := ub.JWT.GenerateToken(usersData.UUID, usersData.Email, usersData.Role)
	if err != nil {
		return "", errors.New("failed generate token")
	}
	return token, nil
}

// Register implements users_domain.Business
func (ub UsersBusiness) Register(domain users_domain.Users, file interface{}) error {
	// make uuid
	uuidData := uuid.New()
	domain.UUID = uuidData.String()
	if domain.Email == "admin@bayeue.com" {
		domain.Role = "admin"
	} else {
		domain.Role = "customer"
	}
	// upload image
	img, _ := claudinary.ImageUploadHelper(file, "users")

	domain.Image = img
	if domain.Image == "" {
		domain.Image = "https://res.cloudinary.com/dt91kxctr/image/upload/v1655825545/go-bayeue/users/download_o1yrxx.png"
	}

	fmt.Println("image ", domain.Image)

	// store data
	err := ub.UsersRepo.Store(domain)
	if err != nil {
		return err
	}
	return nil
}

func NewUsersHandler(userRepo users_domain.Repo, jwt middlewares.ConfigJwt) users_domain.Business {
	return UsersBusiness{
		JWT:       jwt,
		UsersRepo: userRepo,
	}
}
