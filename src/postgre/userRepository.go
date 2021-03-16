package postgre

import (
	"wordy/src/user"

	"gorm.io/gorm"
)

// UserRepository gorm implementation
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository constructor
func NewUserRepository(DB *gorm.DB) user.Repository {
	return UserRepository{
		DB: DB,
	}
}

// Save User
func (r UserRepository) Save(user user.User) user.User {
	r.DB.Save(&user)
	return user
}
