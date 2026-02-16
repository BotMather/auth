package auth

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string     `gorm:"first_name"`
	LastName    string     `gorm:"last_name"`
	Phone       string     `gorm:"phone;unique"`
	Password    string     `gorm:"password"`
	ValidatedAT *time.Time `gorm:"validated_at"`
	Role        string     `gorm:"role;default:user"`
}

func (*User) TableName() string {
	return "users"
}

type Otp struct {
	gorm.Model
	Phone string `gorm:"phone;unique"`
	Code  string `gorm:"code"`
	Exp   string `gorm:"exp"`
}

func (*Otp) TableName() string {
	return "otp"
}
