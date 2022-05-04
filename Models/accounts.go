package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model        // converting this struct into object model via ORM
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	//Email     string `json:"email"`
	//Addr      string `json:"address"`
}
