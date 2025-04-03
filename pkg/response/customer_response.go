package response

import (
	"github.com/DoHongKien/go-structure/internal/models"
)

type CustomerResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Active   string `json:"active"`
}

func MapCustomer(customer models.Customer) *CustomerResponse {
	return &CustomerResponse{
		ID:       customer.ID,
		FullName: customer.FullName,
		Email:    customer.Email,
		Phone:    customer.Phone,
		Active:   customer.Status,
	}
}
