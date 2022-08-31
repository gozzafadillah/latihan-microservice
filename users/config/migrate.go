package users_config

import (
	users_mysql "gozzafadillah/users/repository/mysql"

	"gorm.io/gorm"
)

func AutoMigrate(DB *gorm.DB) {

	DB.AutoMigrate(&users_mysql.Users{})

}
