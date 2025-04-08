package repo

import (
	"github.com/DoHongKien/go-structure/global"
	"github.com/DoHongKien/go-structure/internal/model"
	"github.com/DoHongKien/go-structure/internal/model/dto"
	"gorm.io/gorm"
)

// ICustomerRepository defines the contract for customer data operations
type ICustomerRepository interface {
	GetAllCustomers(limit, offset int) ([]model.Customer, error)
	GetCustomerByID(id int) (*model.Customer, error)
	GetRawQueryCustomer(id int) (*dto.CustomerRaw, error)
}

type customerRepository struct {
	db *gorm.DB
}

// NewCustomerRepository creates a new customer repository instance
func NewCustomerRepository() ICustomerRepository {
	return &customerRepository{
		db: global.Mdb,
	}
}

// GetAllCustomers retrieves customers with pagination
func (r *customerRepository) GetAllCustomers(limit, offset int) ([]model.Customer, error) {
	var customers []model.Customer

	result := r.db.
		Limit(limit).
		Offset(offset).
		Find(&customers)

	if result.Error != nil {
		return nil, result.Error
	}

	return customers, nil
}

// GetCustomerByID retrieves a single customer by ID
func (r *customerRepository) GetCustomerByID(id int) (*model.Customer, error) {
	var customer model.Customer

	if err := r.db.First(&customer, id).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

// GetRawQueryCustomer retrieves specific customer fields using raw SQL
func (r *customerRepository) GetRawQueryCustomer(id int) (*dto.CustomerRaw, error) {
	var customerData dto.CustomerRaw

	query := "SELECT username, email FROM customer WHERE id = ?"
	if err := r.db.Raw(query, id).Scan(&customerData).Error; err != nil {
		return nil, err
	}

	return &customerData, nil
}
