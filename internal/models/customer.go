package models

import "time"

type Customer struct {
	ID        int       `gorm:"primaryKey;column:id"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone"`
	Address   string    `gorm:"column:address"`
	Active    bool      `gorm:"column:active;default:true"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
