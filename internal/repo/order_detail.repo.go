package repo

import (
	"github.com/DoHongKien/go-structure/internal/models"
	"gorm.io/gorm"
)

func SaveOrderDetail(db *gorm.DB, orderDetail *models.OrderDetail) (*models.OrderDetail, error) {
	err := db.Create(orderDetail).Error

	if err != nil {
		return nil, err
	}

	return orderDetail, nil
}

func GetOrderDetail(db *gorm.DB, id int) (*models.OrderDetail, error) {
	orderDetail := &models.OrderDetail{}
	err := db.First(orderDetail, id).Error

	if err != nil {
		return nil, err
	}

	return orderDetail, nil
}

func GetAllOrderDetails(db *gorm.DB) ([]models.OrderDetail, error) {
	var orderDetails []models.OrderDetail
	err := db.Preload("Order").Find(&orderDetails).Error

	if err != nil {
		return nil, err
	}

	return orderDetails, nil
}
