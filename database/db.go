package database

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"pismo/account"
	"pismo/transaction"
)

var (
	RNFError = errors.New("record not found")
)

func ConnectDatabase() (*gorm.DB, error) {
	DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	if err = DB.AutoMigrate(&account.Account{}, &transaction.Transaction{}); err != nil {
		return nil, err
	}

	return DB, nil
}
