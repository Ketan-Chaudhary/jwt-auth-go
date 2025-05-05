package database

import (
	"log"
	"myjwtauthentication/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connection, err := gorm.Open(mysql.Open("ketan:1234@/go_auth"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = connection
	connection.AutoMigrate(&model.User{})
}
