package transaction

import (
	"time"
)

type Transaction struct {
	ID          int64
	Amount      int64
	Type        string
	Description string
	CreatedAt   time.Time
	Customer    int64
}

func NewTransaction(id, amount int64, t, description string, createdAt time.Time, customer int64) Transaction {
	return Transaction{
		ID:          id,
		Amount:      amount,
		Type:        t,
		Description: description,
		CreatedAt:   createdAt,
		Customer:    customer,
	}
}
