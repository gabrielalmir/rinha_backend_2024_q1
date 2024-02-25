package dto

import (
	"time"

	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/customer/entity"
)

type CustomerDTO struct {
	Total         int64            `json:"total"`
	StatementDate time.Time        `json:"data_extrato"`
	Transactions  []TransactionDTO `json:"ultimas_transacoes"`
}

func NewCustomerDTO(c *entity.Customer) CustomerDTO {
	var transactionsDTO []TransactionDTO
	for _, t := range c.Transactions {
		transactionsDTO = append(transactionsDTO, NewTransactionDTO(&t))
	}

	return CustomerDTO{
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
