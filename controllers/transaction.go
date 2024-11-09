package controllers

import (
	"encoding/json"
	"net/http"

	"pismo/models"
	"pismo/services"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var input struct {
		AccountID       uint                 `json:"account_id"`
		OperationTypeID models.OperationType `json:"operation_type_id"`
		Amount          float64              `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Adjust the amount based on the operation type
	if input.OperationTypeID != models.CreditVoucher {
		input.Amount = -input.Amount
	}

	transaction := models.Transaction{
		AccountID:       input.AccountID,
		OperationTypeID: input.OperationTypeID,
		Amount:          input.Amount,
	}

	if err := services.CreateTransaction(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
