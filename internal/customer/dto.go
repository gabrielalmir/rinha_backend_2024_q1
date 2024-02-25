package customer

import (
	"time"

	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/transaction"
)

type CustomerDTO struct {
	Total         int64                        `json:"total"`
	StatementDate time.Time                    `json:"data_extrato"`
	Transactions  []transaction.TransactionDTO `json:"ultimas_transacoes"`
}

func NewCustomerDTO(total int64, transactions []transaction.TransactionDTO) CustomerDTO {
	return CustomerDTO{
		Total:         total,
		StatementDate: time.Now(),
		Transactions:  transactions,
	}
}
