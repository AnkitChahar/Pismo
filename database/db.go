package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"pismo/models"
)

func ConnectDatabase(dsn string) (*gorm.DB, error) {
	DB, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	if err = DB.AutoMigrate(&models.Account{}, &models.Transaction{}); err != nil {
		return nil, err
	}

	return DB, nil
}
