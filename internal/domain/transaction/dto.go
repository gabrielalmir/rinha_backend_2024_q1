package transaction

type TransactionDTO struct {
	Description string `json:"descricao"`
	Type        string `json:"tipo"`
	Amount      int64  `json:"valor"`
}

func NewTransactionDTO(t *Transaction) TransactionDTO {
	return TransactionDTO{
		Description: t.Description,
		Type:        t.Type,
		Amount:      t.Amount,
	}
}
