package service

import (
	"github.com/DoHongKien/go-structure/internal/model"
	"github.com/DoHongKien/go-structure/internal/model/dto"
	"github.com/DoHongKien/go-structure/internal/repo"
)

// IOrderDetailService defines operations for order detail service
type IOrderDetailService interface {
	SaveOrderDetail(orderDetail *model.OrderDetail) (*model.OrderDetail, error)
	GetOrderDetail(id int) (*model.OrderDetail, error)
	GetAllOrderDetails() ([]model.OrderDetail, error)
	OrderDetailJoin() ([]dto.OrderJoin, error)
}

// orderDetailService implements IOrderDetailService
type orderDetailService struct {
	repository repo.IOrderDetailRepository
}

// NewOrderDetailService creates a new order detail service instance
func NewOrderDetailService(orderDetailRepo repo.IOrderDetailRepository) IOrderDetailService {
	return &orderDetailService{
		repository: orderDetailRepo,
	}
}

// SaveOrderDetail persists an order detail to storage
func (o *orderDetailService) SaveOrderDetail(orderDetail *model.OrderDetail) (*model.OrderDetail, error) {
	return o.repository.SaveOrderDetail(orderDetail)
}

// GetOrderDetail retrieves an order detail by ID
func (o *orderDetailService) GetOrderDetail(id int) (*model.OrderDetail, error) {
	return o.repository.GetOrderDetail(id)
}

// GetAllOrderDetails retrieves all order details
func (o *orderDetailService) GetAllOrderDetails() ([]model.OrderDetail, error) {
	return o.repository.GetAllOrderDetails()
}

// OrderDetailJoin retrieves order details with joined data
func (o *orderDetailService) OrderDetailJoin() ([]dto.OrderJoin, error) {
	return o.repository.OrderDetailJoin()
}
