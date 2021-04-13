package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:13572468@tcp(tidb.142c4f0b.38bdef7b.us-west-2.prod.aws.tidbcloud.com:4000)/test"), &gorm.Config{})

	if err != nil {
		panic("cannot connect database")
	}
}
