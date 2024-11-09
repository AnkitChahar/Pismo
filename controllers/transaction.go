package controllers

import (
	"encoding/json"
	"net/http"

	"pismo/transaction"
)

type TransactionController struct {
	transactionSvc transaction.TransactionService
}

func NewTransactionController(transactionSvc transaction.TransactionService) *TransactionController {
	return &TransactionController{
		transactionSvc: transactionSvc,
	}
}

func (c *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var input struct {
		AccountID       uint                      `json:"account_id"`
		OperationTypeID transaction.OperationType `json:"operation_type_id"`
		Amount          float64                   `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	txn, err := c.transactionSvc.CreateTransaction(&transaction.Transaction{
		AccountID:       input.AccountID,
		OperationTypeID: input.OperationTypeID,
		Amount:          input.Amount,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(txn); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
