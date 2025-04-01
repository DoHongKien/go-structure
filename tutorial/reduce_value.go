package tutorial

import (
	"fmt"

	m "github.com/DoHongKien/go-structure/internal/models"
	res "github.com/DoHongKien/go-structure/response"
)

func MapArray() {

	arr := []m.Customer{
		{ID: 1, FirstName: "Kien", LastName: "Do", Email: "dohongkien@gmail.com", Phone: "0123456789", Address: "Hanoi", Active: true},
		{ID: 2, FirstName: "John", LastName: "Smith", Email: "johnsmith@gmail.com", Phone: "0987654321", Address: "New York", Active: true},
		{ID: 3, FirstName: "Jane", LastName: "Doe", Email: "janedoe@gmail.com", Phone: "0112233445", Address: "London", Active: false},
		{ID: 4, FirstName: "Alice", LastName: "Brown", Email: "alicebrown@gmail.com", Phone: "0223344556", Address: "Paris", Active: true},
		{ID: 5, FirstName: "Bob", LastName: "White", Email: "bobwhite@gmail.com", Phone: "0334455667", Address: "Berlin", Active: false},
		{ID: 6, FirstName: "Charlie", LastName: "Black", Email: "charlieblack@gmail.com", Phone: "0445566778", Address: "Tokyo", Active: true},
		{ID: 7, FirstName: "Daisy", LastName: "Green", Email: "daisygreen@gmail.com", Phone: "0556677889", Address: "Sydney", Active: true},
		{ID: 8, FirstName: "Eve", LastName: "Blue", Email: "eveblue@gmail.com", Phone: "0667788990", Address: "Toronto", Active: false},
		{ID: 9, FirstName: "Frank", LastName: "Gray", Email: "frankgray@gmail.com", Phone: "0778899001", Address: "Dubai", Active: true},
		{ID: 10, FirstName: "Grace", LastName: "Yellow", Email: "graceyellow@gmail.com", Phone: "0889900112", Address: "Singapore", Active: false},
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
