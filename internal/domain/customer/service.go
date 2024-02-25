package customer

import (
	"gorm.io/gorm"
)

type CustomerService struct {
	db *gorm.DB
}

func NewCustomerService(db *gorm.DB) *CustomerService {
	return &CustomerService{db: db}
}

func (s *CustomerService) GetCustomerStatement(id string) (CustomerDTO, error) {
	return CustomerDTO{}, nil
}
