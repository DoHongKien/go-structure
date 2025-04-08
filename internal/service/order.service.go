package service

import (
	"github.com/DoHongKien/go-structure/internal/model"
	"github.com/DoHongKien/go-structure/internal/repo"
)

// IOrderService defines operations for handling orders
type IOrderService interface {
	CreateOrder(order *model.Order) (*model.Order, error)
	GetOrderByID(id int) (*model.Order, error)
	GetAllOrders() ([]model.Order, error)
}

type orderService struct {
	repo repo.IOrderRepository
}

// NewOrderService creates a new order service instance
func NewOrderService(repo repo.IOrderRepository) IOrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(order *model.Order) (*model.Order, error) {
	return s.repo.SaveOrder(order)
}

func (s *orderService) GetOrderByID(id int) (*model.Order, error) {
	return s.repo.GetOrderByID(id)
}

func (s *orderService) GetAllOrders() ([]model.Order, error) {
	return s.repo.GetAllOrders()
}
