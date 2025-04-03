package service

import (
	"github.com/DoHongKien/go-structure/internal/models"
	"github.com/DoHongKien/go-structure/internal/repo"
)

type IOrderService interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetOrderByID(id int) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
}

type orderService struct {
	orderRepository repo.IOrderRepository
}

func NewOrderService(orderRepository repo.IOrderRepository) IOrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (os *orderService) CreateOrder(order *models.Order) (*models.Order, error) {
	return os.orderRepository.SaveOrder(order)
}

func (os *orderService) GetOrderByID(id int) (*models.Order, error) {
	return os.orderRepository.GetOrderByID(id)
}

func (os *orderService) GetAllOrders() ([]models.Order, error) {
	return os.orderRepository.GetAllOrders()
}
