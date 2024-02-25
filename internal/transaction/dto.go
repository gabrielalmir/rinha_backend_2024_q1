package transaction

type TransactionDTO struct {
	Description string `json:"descricao"`
	Type        string `json:"tipo"`
	Amount      int64  `json:"valor"`
}

func NewTransactionDTO(description string, t string, amount int64) TransactionDTO {
	return TransactionDTO{
		Description: description,
		Type:        t,
		Amount:      amount,
	}
}

func (t *TransactionDTO) GetSignedAmount() int64 {
	if t.Type == "c" {
		return t.Amount
	}
	return -t.Amount
}

func (t *TransactionDTO) IsValid(c Customer) bool {
	return c.GetLimit() >= c.GetBalance()+t.GetSignedAmount()
}
