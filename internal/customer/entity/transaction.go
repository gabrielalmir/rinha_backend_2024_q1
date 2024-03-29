package entity

import "time"

type Transaction struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement;column:id;index"`
	Amount      int64     `gorm:"column:valor;not null"`
	Type        string    `gorm:"column:tipo;not null"`
	Description string    `gorm:"column:descricao;not null"`
	CreatedAt   time.Time `gorm:"column:criado_em;not null"`
	CustomerID  uint64    `gorm:"column:cliente_id;not null;index"`
}

func (t *Transaction) TableName() string {
	return "transacoes"
}
