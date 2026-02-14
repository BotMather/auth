package auth

type User struct {
	FirstName string `gorm:"first_name"`
	LastName  string `gorm:"last_name"`
	Email     string `gorm:"email"`
	Phone     string `gorm:"phone"`
}

