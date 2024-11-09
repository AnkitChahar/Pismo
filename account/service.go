package account

import (
	"gorm.io/gorm"
)

type AccountService interface {
	CreateAccount(account *Account) error
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

func (s *Service) CreateAccount(account *Account) error {
	return s.db.Create(account).Error
}

func (s *Service) GetAccountByID(id uint) (*Account, error) {
	var account Account
	if err := s.db.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}
