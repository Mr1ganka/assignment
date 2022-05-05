package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "assignment/Models"
)

var DB *gorm.DB
var err error

const db_url = "root:root@tcp(localhost:3306)/assignment?charset=utf8mb4&parseTime=True&loc=Local"

//URL format  - "user:password@connectiontype(connectionurl:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

func InitialMigration() {

	DB, err = gorm.Open(mysql.Open(db_url), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB") // exception that arises as runtime
	}
	DB.Migrator().DropTable(&model.Account{}, &model.Address{}, &model.Contact{})
	DB.AutoMigrate(&model.Account{}, &model.Address{}, &model.Contact{})
}
