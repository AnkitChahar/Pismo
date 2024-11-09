package models

type Account struct {
	ID             uint   `json:"account_id" gorm:"primaryKey"`
	DocumentNumber string `json:"document_number" gorm:"unique;not null"`
}
