package model

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() error {
	db, err := gorm.Open(sqlite.Open("data/test.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}
	fmt.Println("Successfully configured sqlite database")
	db.AutoMigrate(&ExchangeRate{})
	DB = db
	return nil
}
