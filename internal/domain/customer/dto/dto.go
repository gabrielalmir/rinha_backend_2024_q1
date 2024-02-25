package dto

import (
	"time"

	customer "github.com/gabrielalmir/rinha_backend_2024_q1/internal/domain/customer/entity"
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/domain/transaction/dto"
	transaction "github.com/gabrielalmir/rinha_backend_2024_q1/internal/domain/transaction/entity"
)

type CustomerDTO struct {
	Total         int64                `json:"total"`
	StatementDate time.Time            `json:"data_extrato"`
	Transactions  []dto.TransactionDTO `json:"ultimas_transacoes"`
}

func NewCustomerDTO(c *customer.Customer, transactions []transaction.Transaction) CustomerDTO {
	var transactionsDTO []dto.TransactionDTO
	for _, t := range transactions {
		transactionsDTO = append(transactionsDTO, dto.NewTransactionDTO(&t))
	}

	return CustomerDTO{
		Total:         c.Balance,
		StatementDate: time.Now(),
		Transactions:  transactionsDTO,
	}
}
