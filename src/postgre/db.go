package postgre

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDBConnection constructor
func NewDBConnection(dsn string) (*gorm.DB, error) {
	// dsn := "host=localhost user=wordy password=123456 dbname=wordy port=5433 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to the database")
	}
	// TODO: Проверить миграцию
	// db.AutoMigrate(&user.User{})

	return db, nil
}
