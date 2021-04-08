package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	NewConfig()
	d := viper.Get("database_host")
	d1 := viper.Get("database_host1")

	fmt.Println(d, d1)
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

func NewConfig() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	err = viper.MergeInConfig() // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}

	if _, err := os.Stat("./.env.local"); err == nil {
		viper.SetConfigName(".env.local")
		err = viper.MergeInConfig() // Find and read the config file
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s ", err))
		}
	}
}
