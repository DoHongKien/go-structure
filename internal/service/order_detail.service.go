package service

import (
	"github.com/DoHongKien/go-structure/internal/models"
	"github.com/DoHongKien/go-structure/internal/models/dto"
	"github.com/DoHongKien/go-structure/internal/repo"
)

type IOrderDetailService interface {
	SaveOrderDetail(orderDetail *models.OrderDetail) (*models.OrderDetail, error)
	GetOrderDetail(id int) (*models.OrderDetail, error)
	GetAllOrderDetails() ([]models.OrderDetail, error)
	OrderDetailJoin() ([]dto.OrderJoin, error)
}

type orderDetailService struct {
	repository repo.IOrderDetailRepository
}

func NewOrderDetailService(orderDetailRepo repo.IOrderDetailRepository) IOrderDetailService {
	return &orderDetailService{
		repository: orderDetailRepo,
	}
}

// GetAllOrderDetails implements IOrderDetailService.
func (o *orderDetailService) GetAllOrderDetails() ([]models.OrderDetail, error) {
	return o.repository.GetAllOrderDetails()
}

// GetOrderDetail implements IOrderDetailService.
func (o *orderDetailService) GetOrderDetail(id int) (*models.OrderDetail, error) {
	return o.repository.GetOrderDetail(id)
}

// OrderDetailJoin implements IOrderDetailService.
func (o *orderDetailService) OrderDetailJoin() ([]dto.OrderJoin, error) {
	return o.repository.OrderDetailJoin()
}

// SaveOrderDetail implements IOrderDetailService.
func (o *orderDetailService) SaveOrderDetail(orderDetail *models.OrderDetail) (*models.OrderDetail, error) {
	return o.repository.SaveOrderDetail(orderDetail)
}
