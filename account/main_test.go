package account

import (
	"log"
	"os"
	"testing"

	"gorm.io/gorm"
	"pismo/database"
	"pismo/models"
)

var (
	testDB *gorm.DB
)

func TestMain(m *testing.M) {
	testDBDSN := "test.db"

	// Remove DB file if it already exists to create a fresh test
	if _, err := os.Stat(testDBDSN); err == nil {
		err = os.Remove(testDBDSN)
		if err != nil {
			log.Fatal(err)
		}
	}

	db, err := database.ConnectDatabase(testDBDSN)
	if err != nil {
		log.Fatal(err)
	}

	testDB = db

	defer os.Remove(testDBDSN)

	m.Run()
}

func insertMockData(accounts ...*models.Account) {
	for _, account := range accounts {
		testDB.Create(account)
	}
}
