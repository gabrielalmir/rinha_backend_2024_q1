package service

import (
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/domain/customer/dto"
	"gorm.io/gorm"
)

type CustomerService struct {
	db *gorm.DB
}

func NewCustomerService(db *gorm.DB) *CustomerService {
	return &CustomerService{db: db}
}

func (s *CustomerService) GetCustomerStatement(id string) (dto.CustomerDTO, error) {
	return dto.CustomerDTO{}, nil
}
