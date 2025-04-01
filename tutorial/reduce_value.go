package tutorial

import (
	"fmt"

	m "github.com/DoHongKien/go-structure/internal/models"
	res "github.com/DoHongKien/go-structure/response"
)

func MapArray() {

	arr := []m.Customer{
		{ID: 1, FullName: "Kien Do", Email: "dohongkien@gmail.com", Phone: "0123456789", Status: "ACTIVE"},
		{ID: 2, FullName: "John Smith", Email: "johnsmith@gmail.com", Phone: "0987654321", Status: "ACTIVE"},
		{ID: 3, FullName: "Jane Doe", Email: "janedoe@gmail.com", Phone: "0112233445", Status: "INACTIVE"},
		{ID: 4, FullName: "Alice Brown", Email: "alicebrown@gmail.com", Phone: "0223344556", Status: "ACTIVE"},
		{ID: 5, FullName: "Bob White", Email: "bobwhite@gmail.com", Phone: "0334455667", Status: "INACTIVE"},
		{ID: 6, FullName: "Charlie Black", Email: "charlieblack@gmail.com", Phone: "0445566778", Status: "ACTIVE"},
		{ID: 7, FullName: "Daisy Green", Email: "daisygreen@gmail.com", Phone: "0556677889", Status: "ACTIVE"},
		{ID: 8, FullName: "Eve Blue", Email: "eveblue@gmail.com", Phone: "0667788990", Status: "INACTIVE"},
		{ID: 9, FullName: "Frank Gray", Email: "frankgray@gmail.com", Phone: "0778899001", Status: "ACTIVE"},
		{ID: 10, FullName: "Grace Yellow", Email: "graceyellow@gmail.com", Phone: "0889900112", Status: "INACTIVE"},
	}

	customerResponse := []res.CustomerResponse{}
	for i := range arr {
		customerResponse = append(customerResponse, *res.MapCustomer(arr[i]))
		// customer := response.MapCustomer(arr[i])
		// fmt.Printf("ID: %v | Name: %v | Email: %v | Phone: %v | Address: %v | Active: %v\n", customer.ID, customer.FullName, customer.Email, customer.Phone, customer.Address, customer.Active)
	}

	for customer := range customerResponse {
		fmt.Printf("ID: %v | Name: %v | Email: %v | Phone: %v | Address: %v | Active: %v\n", customerResponse[customer].ID, customerResponse[customer].FullName, customerResponse[customer].Email, customerResponse[customer].Phone, customerResponse[customer].Address, customerResponse[customer].Active)
	}
}
