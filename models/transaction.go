package models

import "time"

type Transaction struct {
	ID              uint          `json:"transaction_id" gorm:"primaryKey"`
	AccountID       uint          `json:"account_id"`
	OperationTypeID OperationType `json:"operation_type_id"`
	Amount          float64       `json:"amount"`
	EventDate       time.Time     `json:"event_date"`
}
