package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"pismo/models"
)

// ConnectDatabase connects to a SQLite database
func ConnectDatabase() (*gorm.DB, error) {
	DB, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	if err = DB.AutoMigrate(&models.Account{}, &models.Transaction{}); err != nil {
		return nil, err
	}

	return DB, nil
}

// ConnectTestDatabase connects to an in-memory SQLite database for testing purposes
func ConnectTestDatabase() (*gorm.DB, error) {
	DB, err := gorm.Open(sqlite.Open(":memory:?cache=shared&mode=memory"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	if err = DB.AutoMigrate(&models.Account{}, &models.Transaction{}); err != nil {
		return nil, err
	}

	return DB, nil
}
