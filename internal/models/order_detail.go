package models

type OrderDetail struct {
	ID          int     `gorm:"primaryKey;column:id"`
	Price       float64 `gorm:"column:price"`
	ProductName string  `gorm:"column:product_name"`
	Quantity    int     `gorm:"column:quantity"`
	OrderID     int     `gorm:"column:order_id"`
	Order       Order   `gorm:"foreignKey:OrderID"`
}

func (OrderDetail) TableName() string {
	return "order_detail"
}
