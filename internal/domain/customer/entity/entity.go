package entity

import "github.com/gabrielalmir/rinha_backend_2024_q1/internal/domain/transaction/entity"

type Customer struct {
	ID           int64                `gorm:"primaryKey;autoIncrement;column:id"`
	Limit        int64                `gorm:"column:limite;not null"`
	Balance      int64                `gorm:"column:saldo;not null"`
	Transactions []entity.Transaction `gorm:"foreignKey:CustomerID"`
}

func NewCustomer(limit, balance int64) *Customer {
	return &Customer{
		Limit:   limit,
		Balance: balance,
	}
}
