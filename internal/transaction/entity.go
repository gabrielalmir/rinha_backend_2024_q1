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
	Customer    Customer
}

type Customer interface {
	GetID() int64
	GetLimit() int64
	GetBalance() int64
	GetTransactions() []Transaction
}

func NewTransaction(id, amount int64, t, description string, createdAt time.Time, customer Customer) Transaction {
	return Transaction{
		ID:          id,
		Amount:      amount,
		Type:        t,
		Description: description,
		CreatedAt:   createdAt,
		Customer:    customer,
	}
}

func (t *Transaction) ToDTO() TransactionDTO {
	return NewTransactionDTO(t.Description, t.Type, t.Amount)
}

func (t *Transaction) GetID() int64 {
	return t.ID
}

func (t *Transaction) GetAmount() int64 {
	return t.Amount
}

func (t *Transaction) GetType() string {
	return t.Type
}

func (t *Transaction) GetDescription() string {
	return t.Description
}

func (t *Transaction) GetCreatedAt() time.Time {
	return t.CreatedAt
}

func (t *Transaction) GetCustomer() Customer {
	return t.Customer
}
