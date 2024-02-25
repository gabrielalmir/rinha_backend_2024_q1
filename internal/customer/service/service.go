package service

import (
	"errors"

	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/customer/dto"
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/customer/entity"
	"gorm.io/gorm"
)

type CustomerService struct {
	db *gorm.DB
}

func NewCustomerService(db *gorm.DB) *CustomerService {
	return &CustomerService{db: db}
}

func (s *CustomerService) GetCustomerStatement(id string) (dto.CustomerStatementDTO, error) {
	var customer entity.Customer
	result := s.db.Preload("Transactions").First(&customer, id)

	if result.Error != nil {
		return dto.CustomerStatementDTO{}, result.Error
	}

	return dto.NewCustomerStatementDTO(&customer), nil
}

func (s *CustomerService) CreateTransaction(id string, transactionDTO dto.TransactionDTO) (dto.CustomerDTO, error) {
	var customer entity.Customer
	result := s.db.First(&customer, id)

	if result.Error != nil {
		return dto.CustomerDTO{}, result.Error
	}

	if !transactionDTO.IsValid(&customer) {
		return dto.CustomerDTO{}, errors.New("invalid transaction")
	}

	transaction := entity.Transaction{
		Amount:      transactionDTO.Amount,
		Type:        transactionDTO.Type,
		Description: transactionDTO.Description,
		CustomerID:  customer.ID,
	}

	customer.Balance += transactionDTO.SignedAmount()
	s.db.Save(&customer)

	s.db.Create(&transaction)

	return dto.NewCustomerDTO(&customer), nil
}
