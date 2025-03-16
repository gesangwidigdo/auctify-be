package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		panic("Environment variable DB_URL not found!")
	}

	fmt.Println("Connecting to DB with DSN:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) 
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	fmt.Println("Database connected successfully!")
	return db
}
