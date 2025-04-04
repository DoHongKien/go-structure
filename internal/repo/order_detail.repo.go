package repo

import (
	"github.com/DoHongKien/go-structure/global"
	"github.com/DoHongKien/go-structure/internal/models"
	"github.com/DoHongKien/go-structure/internal/models/dto"
	"gorm.io/gorm"
)

// IOrderDetailRepository defines order detail repository operations
type IOrderDetailRepository interface {
	SaveOrderDetail(orderDetail *models.OrderDetail) (*models.OrderDetail, error)
	GetOrderDetail(id int) (*models.OrderDetail, error)
	GetAllOrderDetails() ([]models.OrderDetail, error)
	OrderDetailJoin() ([]dto.OrderJoin, error)
}

// orderDetailRepository implements IOrderDetailRepository
type orderDetailRepository struct {
	db *gorm.DB
}

// NewOrderDetailRepository creates a new order detail repository
func NewOrderDetailRepository() IOrderDetailRepository {
	return &orderDetailRepository{
		db: global.Mdb,
	}
}

// SaveOrderDetail creates a new order detail record
func (o *orderDetailRepository) SaveOrderDetail(orderDetail *models.OrderDetail) (*models.OrderDetail, error) {
	if err := o.db.Create(orderDetail).Error; err != nil {
		return nil, err
	}
	return orderDetail, nil
}

// GetOrderDetail retrieves an order detail by ID
func (o *orderDetailRepository) GetOrderDetail(id int) (*models.OrderDetail, error) {
	orderDetail := &models.OrderDetail{}
	if err := o.db.First(orderDetail, id).Error; err != nil {
		return nil, err
	}
	return orderDetail, nil
}

// GetAllOrderDetails retrieves all order details with their orders
func (o *orderDetailRepository) GetAllOrderDetails() ([]models.OrderDetail, error) {
	var orderDetails []models.OrderDetail
	if err := o.db.Preload("Order").Find(&orderDetails).Error; err != nil {
		return nil, err
	}
	return orderDetails, nil
}

// OrderDetailJoin retrieves joined order details data
func (o *orderDetailRepository) OrderDetailJoin() ([]dto.OrderJoin, error) {
	var orderDetails []dto.OrderJoin
	query := o.db.Table("order_detail od").
		Select("o.id as order_id, o.code, od.price, od.product_name as product_name, od.quantity").
		Joins("JOIN orders o ON o.id = od.order_id")
	
	if err := query.Scan(&orderDetails).Error; err != nil {
		return nil, err
	}
	return orderDetails, nil
}
