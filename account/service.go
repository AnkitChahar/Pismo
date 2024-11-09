package account

import (
	"fmt"

	"gorm.io/gorm"
	"pismo/models"
)

//go:generate mockgen -source=service.go -destination=./mocks/service_mocks.go -package=mocks
type AccountService interface {
	CreateAccount(account *models.Account) (*models.Account, error)
	GetAccountByID(id uint) (*models.Account, error)
}

type Service struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

var _ AccountService = (*Service)(nil)

func (s *Service) CreateAccount(account *models.Account) (*models.Account, error) {
	if account == nil {
		return nil, fmt.Errorf("mandatory args missing")
	}

	res := s.db.Create(account)
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Statement.Model.(*models.Account), nil
}

func (s *Service) GetAccountByID(id uint) (*models.Account, error) {
	if id == 0 {
		return nil, fmt.Errorf("mandatory args missing")
	}

	var account models.Account
	if err := s.db.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}
