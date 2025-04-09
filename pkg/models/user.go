package models

const (
	UserRoleAdmin = 0
	UserRoleUser  = 1
)

type UserRole int

type User struct {
	Name     string `gorm:"unique"`
	Password string
	Email    string
	Role     UserRole
}
