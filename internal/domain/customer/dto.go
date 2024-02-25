package customer

import (
	"time"

	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/domain/transaction"
)

type CustomerDTO struct {
	Total         int64                        `json:"total"`
	StatementDate time.Time                    `json:"data_extrato"`
	Transactions  []transaction.TransactionDTO `json:"ultimas_transacoes"`
}

func NewCustomerDTO(c *Customer, transactions []transaction.Transaction) CustomerDTO {
	var transactionsDTO []transaction.TransactionDTO
	for _, t := range transactions {
		transactionsDTO = append(transactionsDTO, transaction.NewTransactionDTO(&t))
	}

	return CustomerDTO{
		Total:         c.Balance,
		StatementDate: time.Now(),
		Transactions:  transactionsDTO,
	}
}
