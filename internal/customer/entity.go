package customer

import (
	"time"

	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/transaction"
)

type Customer struct {
	ID           int64
	Limit        int64
	Balance      int64
	Transactions []Transaction
}

type Transaction interface {
	GetID() int64
	GetAmount() int64
	GetType() string
	GetDescription() string
	GetCreatedAt() time.Time
	GetCustomer() Customer
	ToDTO() transaction.TransactionDTO
}

func NewCustomer(id, limit, balance int64) Customer {
	return Customer{
		ID:      id,
		Limit:   limit,
		Balance: balance,
	}
}

func (c *Customer) ToDTO() CustomerDTO {
	total := int64(0)
	for _, t := range c.Transactions {
		if t.GetType() == "c" {
			total += t.GetAmount()
		} else {
			total -= t.GetAmount()
		}
	}

	transactions := make([]transaction.TransactionDTO, len(c.Transactions))
	for i, t := range c.Transactions {
		transactions[i] = t.ToDTO()
	}

	return NewCustomerDTO(total, transactions)
}

func (c *Customer) GetID() int64 {
	return c.ID
}

func (c *Customer) GetLimit() int64 {
	return c.Limit
}

func (c *Customer) GetBalance() int64 {
	return c.Balance
}

func (c *Customer) GetTransactions() []Transaction {
	return c.Transactions
}
