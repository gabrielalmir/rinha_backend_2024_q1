package seeder

import (
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/customer/entity"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	var count int64
	db.Model(&entity.Customer{}).Count(&count)

	if count > 0 {
		return nil
	}

	customers := []entity.Customer{
		*entity.NewCustomer(100000, 0),
		*entity.NewCustomer(80000, 0),
		*entity.NewCustomer(1000000, 0),
		*entity.NewCustomer(10000000, 0),
		*entity.NewCustomer(500000, 0),
	}

	for _, customer := range customers {
		if err := db.Create(&customer).Error; err != nil {
			return err
		}
	}

	return nil
}
