package users_config

import (
	// "os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBNAME string
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT string

	JWTSecret string
}

var Conf Config

func InitDB() (DB *gorm.DB) {
	Conf = Config{
		DBNAME:    "DB_Users",
		DBUSER:    "root",
		DBPASS:    "",
		DBHOST:    "localhost",
		DBPORT:    "3306",
		JWTSecret: "latihan-microservice",
	}

	var err error

	dsn := Conf.DBUSER + ":" + Conf.DBPASS + "@tcp(" + Conf.DBHOST + ":" + Conf.DBPORT + ")/" + Conf.DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return
}
