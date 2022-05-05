package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model        // converting this struct into object model via ORM
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`

	AddressId uint    `json:"-"`
	Address   Address `json:"address"`

	ContactId uint    `json:"-"`
	Contact   Contact `json:"contact"`
	//Email     string `json:"email"`
	//Addr      string `json:"address"`
}
