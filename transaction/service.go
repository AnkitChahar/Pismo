package transaction

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"pismo/account"
	"pismo/models"
)

type TransactionService interface {
	CreateTransaction(transaction *models.Transaction) (*models.Transaction, error)
}

type Service struct {
	db         *gorm.DB
	accountSvc account.AccountService
}

func NewService(db *gorm.DB, accountSvc account.AccountService) *Service {
	return &Service{
		db:         db,
		accountSvc: accountSvc,
	}
}

var _ TransactionService = (*Service)(nil)

func (s *Service) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	accountId := transaction.AccountID

	// Validate if account exists with the given accountId
	_, errAccount := s.accountSvc.GetAccountByID(accountId)
	if errAccount != nil {
		return nil, fmt.Errorf("error while getting account: %w", errAccount)
	}

	transaction.EventDate = time.Now()
	res := s.db.Create(transaction)
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Statement.Model.(*models.Transaction), nil
}
