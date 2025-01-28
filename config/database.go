package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	var db *gorm.DB
	dsn := os.Getenv("DB_URL")
	var err error
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic("Failed to connect to database")
	}

	return db
}
