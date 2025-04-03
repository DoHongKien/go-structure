package repo

import (
	"github.com/DoHongKien/go-structure/internal/models"
	"github.com/DoHongKien/go-structure/internal/models/dto"
	"gorm.io/gorm"
)

type ICustomerRepository interface {
	GetAllCustomers(limit, offset int) ([]models.Customer, error)
	GetCustomerByID(id int) (*models.Customer, error)
	GetRawQueryCustomer(id int) (*dto.CustomerRaw, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) ICustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) GetAllCustomers(limit, offset int) ([]models.Customer, error) {
	var customers []models.Customer
	err := r.db.Limit(limit).Offset(offset).Find(&customers).Error
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *customerRepository) GetCustomerByID(id int) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *customerRepository) GetRawQueryCustomer(id int) (*dto.CustomerRaw, error) {
	var customerRaws dto.CustomerRaw
	err := r.db.Raw("SELECT username, email FROM customer WHERE id = ?", id).Scan(&customerRaws).Error
	if err != nil {
		return nil, err
	}

	return &customerRaws, nil
}
