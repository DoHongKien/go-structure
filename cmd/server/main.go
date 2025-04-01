package main

import (
	"fmt"

	"github.com/DoHongKien/go-structure/configs"
	"github.com/DoHongKien/go-structure/internal/repo"
)

func main() {

	// Initialize MySQL connection
	mysqlDB := configs.InitMySQL()

	orderDetails, err := repo.GetAllOrderDetails(mysqlDB)

	if err != nil {
		fmt.Println("Error fetching orderDetails:", err)
		return
	}

	for _, orderDetail := range orderDetails {
		fmt.Printf("OrderDetailId: %d | %s | %s\n", orderDetail.ID, "ORDER-"+fmt.Sprint(orderDetail.Order.ID), orderDetail.Order.CreatedAt.Format("2006-01-02 15:04"))
	}

	// r := router.NewRouter()

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
