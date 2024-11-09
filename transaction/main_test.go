package transaction

import (
	"log"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"pismo/account/mocks"
	"pismo/database"
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

type MockDeps struct {
	AccountSvc *mocks.MockAccountService
}

func setupServiceWithMocks(t *testing.T) (*Service, *MockDeps) {
	ctrl := gomock.NewController(t)

	accountSvc := mocks.NewMockAccountService(ctrl)

	return &Service{
			db:         testDB,
			accountSvc: accountSvc,
		}, &MockDeps{
			AccountSvc: accountSvc,
		}
}
