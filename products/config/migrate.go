package product_config

import (
	users_mysql "gozzafadillah/products/repository/mysql"

	"gorm.io/gorm"
)

func AutoMigrate(DB *gorm.DB) {

	DB.AutoMigrate(&users_mysql.Category{}, &users_mysql.Product{}, &users_mysql.Detail{})

}
