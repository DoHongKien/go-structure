package service

import (
	"github.com/DoHongKien/go-structure/internal/models"
	"github.com/DoHongKien/go-structure/internal/models/dto"
	"github.com/DoHongKien/go-structure/internal/repo"
)

type CustomerService struct {
	repo *repo.CustomerRepository
}

func NewCustomerService(repo *repo.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) GetAllCustomers() ([]models.Customer, error) {
	return s.repo.GetAllCustomers()
}

func (s *CustomerService) GetCustomerByID(id int) (*models.Customer, error) {
	return s.repo.GetCustomerByID(id)
}

func (s *CustomerService) GetRawQueryCustomer(id int) (*dto.CustomerRaw, error) {
	return s.repo.GetRawQueryCustomer(id)
}
