package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	NewConfig()
	var dbConnection = viper.GetString("database")
	fmt.Print(dbConnection)
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
