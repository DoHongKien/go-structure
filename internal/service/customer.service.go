package service

import (
	"github.com/DoHongKien/go-structure/internal/model"
	"github.com/DoHongKien/go-structure/internal/model/dto"
	"github.com/DoHongKien/go-structure/internal/repo"
)

// ICustomerService defines the interface for customer-related operations
type ICustomerService interface {
	GetAllCustomers(limit, offset int) ([]model.Customer, error)
	GetCustomerByID(id int) (*model.Customer, error)
	GetRawQueryCustomer(id int) (*dto.CustomerRaw, error)
}

// customerService implements the ICustomerService interface
type customerService struct {
	repo repo.ICustomerRepository
}

// NewCustomerService creates a new instance of customerService
func NewCustomerService(repo repo.ICustomerRepository) ICustomerService {
	return &customerService{repo: repo}
}

// GetAllCustomers retrieves all customers with pagination
func (s *customerService) GetAllCustomers(limit, offset int) ([]model.Customer, error) {
	return s.repo.GetAllCustomers(limit, offset)
}

// GetCustomerByID retrieves a customer by their ID
func (s *customerService) GetCustomerByID(id int) (*model.Customer, error) {
	return s.repo.GetCustomerByID(id)
}

// GetRawQueryCustomer retrieves raw customer data by ID
func (s *customerService) GetRawQueryCustomer(id int) (*dto.CustomerRaw, error) {
	return s.repo.GetRawQueryCustomer(id)
}
