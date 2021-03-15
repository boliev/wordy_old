package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// dsn := "host=localhost user=wordy password=123456 dbname=wordy port=5433 sslmode=disable TimeZone=Europe/Moscow"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("Cannot connect to the database")
	// }
	// db.AutoMigrate(&user.User{})

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
