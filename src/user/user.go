package user

import (
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(250)"`
	Email    string `gorm:"type:varchar(250);index:users_email_unique,unique"`
	Password string `gorm:"type:varchar(250)"`
}

// ID user getter
func (u User) ID() uint { return u.Model.ID }

// CreateUser creates a User
func CreateUser(request CreateRequest) User {
	return User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}
