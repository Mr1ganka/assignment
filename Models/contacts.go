package model

// import "gorm.io/gorm"

type Contact struct {
	Id       uint   `gorm:"references:AccountId"`
	Email    string `json:"email"`
	Phone    int64  `json:"phone"`
	Callable bool   `json:"call"`
}
