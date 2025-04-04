package repo

import (
	"fmt"

	"github.com/DoHongKien/go-structure/global"
	"github.com/DoHongKien/go-structure/internal/models"
	"gorm.io/gorm"
)

type IOrderRepository interface {
	SaveOrder(order *models.Order) (*models.Order, error)
	GetOrderByID(id int) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository() IOrderRepository {
	return &orderRepository{db: global.Mdb}
}

func (or *orderRepository) SaveOrder(order *models.Order) (*models.Order, error) {
	err := or.db.Create(order).Error

	if err != nil {
		return nil, err
	}
	return order, nil
}

func (or *orderRepository) GetOrderByID(id int) (*models.Order, error) {
	order := &models.Order{}
	err := or.db.First(order, id).Error

	if err != nil {
		global.Logger.Error(fmt.Sprintf("Error getting order by ID %d: %v", id, err))
		return nil, err
	}

	return order, nil
}

func (or *orderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := or.db.Find(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}
