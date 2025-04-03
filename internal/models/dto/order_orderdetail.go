package dto

type OrderJoin struct {
	OrderID     int    `json:"order_id"`
	Code        string `json:"code"`
	Price       int    `json:"price"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}
