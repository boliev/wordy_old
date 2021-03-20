package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	NewConfig()
	r := gin.Default()
	user, err := InitializeUserController()
	if err != nil {
		panic(err)
	}
	r.POST("/v1/user", func(c *gin.Context) {
		user.Create(c)
	})
	r.Run(":8080")
}
