package users_mysql

import (
	users_domain "gozzafadillah/users/domain"

	"gorm.io/gorm"
)

type UsersRepo struct {
	DB *gorm.DB
}

// GetUserUUID implements users_domain.Repo
func (ur UsersRepo) GetUserUUID(uuid string) (users_domain.Users, error) {
	record := Users{}
	err := ur.DB.Model(&Users{}).Where("uuid = ?", uuid).First(&record).Error
	return ToDomain(record), err
}

// UpdateUser implements users_domain.Repo
func (ur UsersRepo) UpdateUser(domain users_domain.Users, uuid string) error {
	domain.UUID = uuid
	err := ur.DB.Model(&Users{}).Where("uuid = ?", uuid).Updates(domain).Error
	return err
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
