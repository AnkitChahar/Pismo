package database

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"pismo/models"
)

var DB *gorm.DB

var (
	RNFError = errors.New("record not found")
)

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	DB.AutoMigrate(&models.Account{}, &models.Transaction{})
}
