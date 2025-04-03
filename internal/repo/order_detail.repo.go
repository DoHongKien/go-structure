package repo

import (
	"github.com/DoHongKien/go-structure/internal/models"
	"github.com/DoHongKien/go-structure/internal/models/dto"
	"gorm.io/gorm"
)

type IOrderDetailRepository interface {
	SaveOrderDetail(orderDetail *models.OrderDetail) (*models.OrderDetail, error)
	GetOrderDetail(id int) (*models.OrderDetail, error)
	GetAllOrderDetails() ([]models.OrderDetail, error)
	OrderDetailJoin() ([]dto.OrderJoin, error)
}

type orderDetailRepository struct {
	db *gorm.DB
}

func NewOrderDetailRepository(db *gorm.DB) IOrderDetailRepository {
	return &orderDetailRepository{
		db: db,
	}
}

// GetAllOrderDetails implements IOrderDetailRepository.
func (o *orderDetailRepository) GetAllOrderDetails() ([]models.OrderDetail, error) {
	var orderDetails []models.OrderDetail
	err := o.db.Preload("Order").Find(&orderDetails).Error

	if err != nil {
		return nil, err
	}

	return orderDetails, nil
}

// GetOrderDetail implements IOrderDetailRepository.
func (o *orderDetailRepository) GetOrderDetail(id int) (*models.OrderDetail, error) {
	orderDetail := &models.OrderDetail{}
	err := o.db.First(orderDetail, id).Error

	if err != nil {
		return nil, err
	}

	return orderDetail, nil
}

// OrderDetailJoin implements IOrderDetailRepository.
func (o *orderDetailRepository) OrderDetailJoin() ([]dto.OrderJoin, error) {
	var orderDetails []dto.OrderJoin
	err := o.db.Table("order_detail od").Select("o.id as order_id, o.code, od.price, od.product_name as product_name, od.quantity").Joins("JOIN orders o ON o.id = od.order_id").Scan(&orderDetails).Error

	if err != nil {
		return nil, err
	}

	return orderDetails, nil
}

// SaveOrderDetail implements IOrderDetailRepository.
func (o *orderDetailRepository) SaveOrderDetail(orderDetail *models.OrderDetail) (*models.OrderDetail, error) {
	err := o.db.Create(orderDetail).Error

	if err != nil {
		return nil, err
	}

	return orderDetail, nil
}

// func SaveOrderDetail(db *gorm.DB, orderDetail *models.OrderDetail) (*models.OrderDetail, error) {
// 	err := db.Create(orderDetail).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return orderDetail, nil
// }

// func GetOrderDetail(db *gorm.DB, id int) (*models.OrderDetail, error) {
// 	orderDetail := &models.OrderDetail{}
// 	err := db.First(orderDetail, id).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return orderDetail, nil
// }

// func GetAllOrderDetails(db *gorm.DB) ([]models.OrderDetail, error) {
// 	var orderDetails []models.OrderDetail
// 	err := db.Preload("Order").Find(&orderDetails).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return orderDetails, nil
// }

// func OrderDetailJoin(db *gorm.DB) ([]dto.OrderJoin, error) {
// 	var orderDetails []dto.OrderJoin
// 	err := db.Table("order_detail od").Select("o.id as order_id, o.code, od.price, od.product_name as product_name, od.quantity").Joins("JOIN orders o ON o.id = od.order_id").Scan(&orderDetails).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return orderDetails, nil
// }
