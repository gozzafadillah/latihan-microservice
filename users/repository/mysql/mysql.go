package users_mysql

import (
	users_domain "gozzafadillah/users/domain"

	"gorm.io/gorm"
)

type UsersRepo struct {
	DB *gorm.DB
}

// GetUser implements users_domain.Repo
func (ur UsersRepo) GetUser(email string) (users_domain.Users, error) {
	record := Users{}
	err := ur.DB.Where("email = ?", email).First(&record).Error
	return ToDomain(record), err
}

// EmailPasswordCheck implements users_domain.Repo
func (ur UsersRepo) EmailPasswordCheck(email string, password string) error {
	err := ur.DB.Model(&Users{}).Where("email = ? AND password = ?", email, password).Error
	return err
}

// Store implements users_domain.Repo
func (ur UsersRepo) Store(domain users_domain.Users) error {
	err := ur.DB.Create(&domain).Error
	return err
}

func NewUsersRepo(db *gorm.DB) users_domain.Repo {
	return UsersRepo{
		DB: db,
	}
}
