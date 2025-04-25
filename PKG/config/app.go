package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// Connect initializes the database connection
func Connect() {
	dsn := "username:Password@007@tcp(127.0.0.1:Port)/AnimeREST?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	db = d
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
