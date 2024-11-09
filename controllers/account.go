package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"pismo/account"
	"pismo/database"
)

type AccountController struct {
	accountSvc account.AccountService
}

func NewAccountController(accountSvc account.AccountService) *AccountController {
	return &AccountController{
		accountSvc: accountSvc,
	}
}

func (c *AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account account.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.accountSvc.CreateAccount(&account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["accountId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	account, err := c.accountSvc.GetAccountByID(uint(id))
	if err != nil {
		if errors.Is(err, database.RNFError) {
			http.Error(w, "Account not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
