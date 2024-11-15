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

// CreateTransaction creates a new transaction for a account ID, it validates if the account is present, then creates the transaction
func (s *Service) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	accountId := transaction.AccountID

	// Validate if account exists with the given accountId
	_, errAccount := s.accountSvc.GetAccountByID(accountId)
	if errAccount != nil {
		return nil, fmt.Errorf("error while getting account: %w", errAccount)
	}

	if transaction.OperationTypeID != models.CreditVoucher {
		transaction.Amount = -transaction.Amount
	}

	transaction.EventDate = time.Now()
	res := s.db.Create(transaction)
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Statement.Model.(*models.Transaction), nil
}
