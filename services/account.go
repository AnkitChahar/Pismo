package services

import (
	"pismo/database"
	"pismo/models"
)

func CreateAccount(account *models.Account) error {
	return database.DB.Create(account).Error
}

func GetAccountByID(id uint) (*models.Account, error) {
	var account models.Account
	if err := database.DB.First(&account, id).Error; err != nil {
		// check for record not found
		if database.DB.RecordNotFound() {

		}
		return nil, err
	}
	return &account, nil
}
