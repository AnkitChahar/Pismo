package transaction

import (
	"time"

	"gorm.io/gorm"
	"pismo/account"
)

type TransactionService interface {
	CreateTransaction(transaction *Transaction) (*Transaction, error)
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

func (s *Service) CreateTransaction(transaction *Transaction) (*Transaction, error) {
	accountId := transaction.AccountID

	_, errAccount := s.accountSvc.GetAccountByID(accountId)
	if errAccount != nil {
		return nil, errAccount
	}

	transaction.EventDate = time.Now()
	if errCreate := s.db.Create(transaction).Error; errCreate != nil {
		return nil, errCreate
	}

	return transaction, nil
}
