package dto

import "github.com/gabrielalmir/rinha_backend_2024_q1/internal/domain/transaction/entity"

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
