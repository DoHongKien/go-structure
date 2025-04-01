package repo

import (
	"github.com/DoHongKien/go-structure/internal/models"
	"gorm.io/gorm"
)

func SaveOrder(db *gorm.DB, order *models.Order) (*models.Order, error) {
	err := db.Create(order).Error

	if err != nil {
		return nil, err
	}
	return order, nil
}

func GetOrderByID(db *gorm.DB, id int) (*models.Order, error) {
	order := &models.Order{}
	err := db.First(order, id).Error

	if err != nil {
		return nil, err
	}

	return order, nil
}

func GetAllOrders(db *gorm.DB) ([]models.Order, error) {
	var orders []models.Order
	err := db.Find(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}
