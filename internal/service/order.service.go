package service

import (
	"github.com/DoHongKien/go-structure/internal/models"
	"github.com/DoHongKien/go-structure/internal/repo"
)

// IOrderService defines operations for handling orders
type IOrderService interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetOrderByID(id int) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
}

type orderService struct {
	repo repo.IOrderRepository
}

// NewOrderService creates a new order service instance
func NewOrderService(repo repo.IOrderRepository) IOrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(order *models.Order) (*models.Order, error) {
	return s.repo.SaveOrder(order)
}

func (s *orderService) GetOrderByID(id int) (*models.Order, error) {
	return s.repo.GetOrderByID(id)
}

func (s *orderService) GetAllOrders() ([]models.Order, error) {
	return s.repo.GetAllOrders()
}
