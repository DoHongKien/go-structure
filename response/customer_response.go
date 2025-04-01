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
		FullName: customer.LastName + " " + customer.FirstName,
		Email:    customer.Email,
		Phone:    customer.Phone,
		Address:  customer.Address,
		Active:   convertStatus(customer.Active),
	}
}

func convertStatus(active bool) string {
	if active {
		return "Active"
	} else {
		return "InActive"
	}
}
