package models

import (
	"time"
)

type Order struct {
	ID         int       `gorm:"primaryKey;column:id"`
	Code       string    `gorm:"column:code"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	TotalPrice float32   `gorm:"column:total_price"`
}
