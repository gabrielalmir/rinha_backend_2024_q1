package customer

type Customer struct {
	ID      int64
	Limit   int64
	Balance int64
}

func NewCustomer(id, limit, balance int64) *Customer {
	return &Customer{
		ID:      id,
		Limit:   limit,
		Balance: balance,
	}
}
