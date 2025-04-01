package repo

import (
	"github.com/DoHongKien/go-structure/internal/models"
	"gorm.io/gorm"
)

func GetAllCustomers(db *gorm.DB) ([]models.Customer, error) {
	var customers []models.Customer
	err := db.Find(&customers).Error
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func GetCustomerByID(db *gorm.DB, id int) (*models.Customer, error) {
	var customer models.Customer
	err := db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
