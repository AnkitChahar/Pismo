package account

import (
	"fmt"

	"gorm.io/gorm"
)

type AccountService interface {
	CreateAccount(account *Account) (*Account, error)
	GetAccountByID(id uint) (*Account, error)
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

func (s *Service) CreateAccount(account *Account) (*Account, error) {
	if account == nil {
		return nil, fmt.Errorf("mandatory args missing")
	}

	res := s.db.Create(account)
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Statement.Model.(*Account), nil
}

func (s *Service) GetAccountByID(id uint) (*Account, error) {
	if id == 0 {
		return nil, fmt.Errorf("mandatory args missing")
	}

	var account Account
	if err := s.db.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}
