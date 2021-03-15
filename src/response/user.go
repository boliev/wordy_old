package response

import (
	"wordy/src/user"
)

// User response DTO
type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUserFromDomain creates User response DTO from domain User
func CreateUserFromDomain(domainUser user.User) User {
	return User{
		ID:    domainUser.ID(),
		Name:  domainUser.Name,
		Email: domainUser.Email,
	}
}
