package services

import (
	"time"

	"pismo/database"
	"pismo/models"
)

func CreateTransaction(transaction *models.Transaction) error {
	transaction.EventDate = time.Now()
	return database.DB.Create(transaction).Error
}
