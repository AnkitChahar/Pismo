package transaction

import (
	"log"
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
	db, err := database.ConnectTestDatabase()
	if err != nil {
		log.Fatal(err)
	}

	testDB = db

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
