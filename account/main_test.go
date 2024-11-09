package account

import (
	"log"
	"testing"

	"gorm.io/gorm"
	"pismo/database"
	"pismo/models"
)

var (
	testDB *gorm.DB
)

func TestMain(m *testing.M) {
	db, err := database.ConnectTestDatabase()
	if err != nil {
		log.Fatal(err)
	}

	testDB = db

	m.Run()
}

func insertMockData(accounts ...*models.Account) {
	for _, account := range accounts {
		testDB.Create(account)
	}
}
