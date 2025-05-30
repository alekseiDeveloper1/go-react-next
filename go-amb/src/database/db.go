package database

import (
	"amb/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/amb"), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!")
	}

}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
