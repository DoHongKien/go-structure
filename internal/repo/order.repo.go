package repo

import (
	"fmt"

	"github.com/DoHongKien/go-structure/global"
	"github.com/DoHongKien/go-structure/internal/models"
	"gorm.io/gorm"
)

// IOrderRepository defines the contract for order repository operations
type IOrderRepository interface {
	SaveOrder(order *models.Order) (*models.Order, error)
	GetOrderByID(id int) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new instance of order repository
func NewOrderRepository() IOrderRepository {
	return &orderRepository{db: global.Mdb}
}

func (or *orderRepository) SaveOrder(order *models.Order) (*models.Order, error) {
	if err := or.db.Create(order).Error; err != nil {
		global.Logger.Error(fmt.Sprintf("Error creating order: %v", err))
		return nil, err
	}
	return order, nil
}

func (or *orderRepository) GetOrderByID(id int) (*models.Order, error) {
	order := &models.Order{}
	if err := or.db.First(order, id).Error; err != nil {
		global.Logger.Error(fmt.Sprintf("Error getting order by ID %d: %v", id, err))
		return nil, err
	}
	return order, nil
}

func (or *orderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := or.db.Find(&orders).Error; err != nil {
		global.Logger.Error(fmt.Sprintf("Error getting all orders: %v", err))
		return nil, err
	}
	return orders, nil
}
