package entity

type Customer struct {
	ID           uint64        `gorm:"primaryKey;autoIncrement;column:id;index"`
	Limit        int64         `gorm:"column:limite;not null"`
	Balance      int64         `gorm:"column:saldo;not null"`
	Transactions []Transaction `gorm:"foreignKey:CustomerID"`
}

func (c *Customer) TableName() string {
	return "clientes"
}

func NewCustomer(limit, balance int64) *Customer {
	return &Customer{
		Limit:   limit,
		Balance: balance,
	}
}
