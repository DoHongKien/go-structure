package service

import (
	"github.com/DoHongKien/go-structure/internal/models"
	"github.com/DoHongKien/go-structure/internal/models/dto"
	"github.com/DoHongKien/go-structure/internal/repo"
)

type ICustomerService interface {
	GetAllCustomers(limit, offset int) ([]models.Customer, error)
	GetCustomerByID(id int) (*models.Customer, error)
	GetRawQueryCustomer(id int) (*dto.CustomerRaw, error)
}

type customerService struct {
	repo repo.ICustomerRepository
}

func NewCustomerService(repo repo.ICustomerRepository) ICustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) GetAllCustomers(limit, offset int) ([]models.Customer, error) {
	return s.repo.GetAllCustomers(limit, offset)
}

func (s *customerService) GetCustomerByID(id int) (*models.Customer, error) {
	return s.repo.GetCustomerByID(id)
}

func (s *customerService) GetRawQueryCustomer(id int) (*dto.CustomerRaw, error) {
	return s.repo.GetRawQueryCustomer(id)
}
