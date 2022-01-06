package model

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// database object
var DB *gorm.DB

// Establish connectivity to the database
func SetupDatabase(test bool) error {
	path := "data/test.db"
	if test {
		path = "../data/test.db"
	}
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}
	fmt.Println("Successfully configured sqlite database")
	db.AutoMigrate(&ExchangeRate{})
	DB = db
	return nil
}
