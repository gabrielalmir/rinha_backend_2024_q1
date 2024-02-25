package dto

import (
	"time"

	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/customer/entity"
)

type CustomerDTO struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}

func NewCustomerDTO(c *entity.Customer) CustomerDTO {
	return CustomerDTO{
		Limit:   c.Limit,
		Balance: c.Balance,
	}
}

type CustomerStatementDTO struct {
	Total         int64            `json:"total"`
	StatementDate time.Time        `json:"data_extrato"`
	Transactions  []TransactionDTO `json:"ultimas_transacoes"`
}

func NewCustomerStatementDTO(c *entity.Customer) CustomerStatementDTO {
	transactionsDTO := make([]TransactionDTO, 0, len(c.Transactions))

	for _, t := range c.Transactions {
		transactionsDTO = append(transactionsDTO, NewTransactionDTO(&t))
	}

	return CustomerStatementDTO{
		Total:         c.Balance,
		StatementDate: time.Now(),
		Transactions:  transactionsDTO,
	}
}

type TransactionDTO struct {
	Description string `json:"descricao"`
	Type        string `json:"tipo"`
	Amount      int64  `json:"valor"`
}

func NewTransactionDTO(t *entity.Transaction) TransactionDTO {
	return TransactionDTO{
		Description: t.Description,
		Type:        t.Type,
		Amount:      t.Amount,
	}
}

func (t *TransactionDTO) SignedAmount() int64 {
	if t.Type == "c" {
		return t.Amount
	}

	return -t.Amount
}

func (t *TransactionDTO) IsValid(c *entity.Customer) bool {
	return (t.Type == "d" || t.Type == "c") && c.Balance+t.SignedAmount() <= c.Limit
}
