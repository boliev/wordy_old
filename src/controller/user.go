package controller

import (
	"net/http"
	"wordy/src/postgre"
	"wordy/src/request"
	"wordy/src/response"
	"wordy/src/user"

	"github.com/gin-gonic/gin"
)

// User Controller struct
type User struct {
	repository user.Repository
	// repository postgre.UserRepository
}

// CreateUserController Constructor
func CreateUserController(repository postgre.UserRepository) User {
	return User{
		repository: repository,
	}
}

// Create User Action
func (u User) Create(c *gin.Context) {
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateBadRequest("wrong parameters"))
		return
	}

	var createRequest = user.CreateRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	var user = user.CreateUser(createRequest)
	u.repository.Save(user)

	c.JSON(http.StatusOK, response.CreateUserFromDomain(user))

}
