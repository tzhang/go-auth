package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	//_, err := gorm.Open(mysql.Open("root:13572468@tcp(tidb.d8334a59.38bdef7b.us-west1.prod.gcp.tidbcloud.com:4000)/test"), &gorm.Config{})
	_, err := gorm.Open(mysql.Open("root:n1cetest@/go-dev"), &gorm.Config{})

	if err != nil {
		panic("cannot connect database")
	}
}
