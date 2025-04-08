package repo

import (
	"fmt"

	"github.com/DoHongKien/go-structure/global"
	"github.com/DoHongKien/go-structure/internal/model"
	"gorm.io/gorm"
)

// IOrderRepository defines the contract for order repository operations
type IOrderRepository interface {
	SaveOrder(order *model.Order) (*model.Order, error)
	GetOrderByID(id int) (*model.Order, error)
	GetAllOrders() ([]model.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new instance of order repository
func NewOrderRepository() IOrderRepository {
	return &orderRepository{db: global.Mdb}
}

func (or *orderRepository) SaveOrder(order *model.Order) (*model.Order, error) {
	if err := or.db.Create(order).Error; err != nil {
		global.Logger.Error(fmt.Sprintf("Error creating order: %v", err))
		return nil, err
	}
	return order, nil
}

func (or *orderRepository) GetOrderByID(id int) (*model.Order, error) {
	order := &model.Order{}
	if err := or.db.First(order, id).Error; err != nil {
		global.Logger.Error(fmt.Sprintf("Error getting order by ID %d: %v", id, err))
		return nil, err
	}
	return order, nil
}

func (or *orderRepository) GetAllOrders() ([]model.Order, error) {
	var orders []model.Order
	if err := or.db.Find(&orders).Error; err != nil {
		global.Logger.Error(fmt.Sprintf("Error getting all orders: %v", err))
		return nil, err
	}
	return orders, nil
}
