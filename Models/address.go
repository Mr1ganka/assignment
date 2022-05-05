package model

// import (
// 	"gorm.io/gorm"
// )

type Address struct {
	Id        uint   `gorm:"references:AccountId"`
	Street    string `json:"street"`
	Apartment string `json:"apartment"`
	City      string `json:"city"`
	Country   string `json:"country"`
}
