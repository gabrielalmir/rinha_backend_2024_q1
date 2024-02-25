package service

import (
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

func (s *CustomerService) GetCustomerStatement(id string) (dto.CustomerDTO, error) {
	var customer entity.Customer
	result := s.db.First(&customer, id)

	if result.Error != nil {
		return dto.CustomerDTO{}, result.Error
	}

	return dto.NewCustomerDTO(&customer), nil
}
