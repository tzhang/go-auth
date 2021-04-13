package database

import (
	"github.com/tzhang/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//_, err := gorm.Open(mysql.Open("root:13572468@tcp(tidb.d8334a59.38bdef7b.us-west1.prod.gcp.tidbcloud.com:4000)/test"), &gorm.Config{})
	connection, err := gorm.Open(mysql.Open("root:n1cetest@/go-dev"), &gorm.Config{})

	DB = connection

	if err != nil {
		panic("cannot connect database")
	}
	connection.AutoMigrate(&models.User{})
}
