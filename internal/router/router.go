package router

import (
	"github.com/DoHongKien/go-structure/internal/controller"
	dbinit "github.com/DoHongKien/go-structure/internal/init"
	"github.com/DoHongKien/go-structure/internal/repo"
	"github.com/DoHongKien/go-structure/internal/service"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	r := gin.Default()

	mysqlDB := dbinit.InitMySQL()

	// Initialize Customer
	customerRepo := repo.NewCustomerRepository(mysqlDB)
	customerService := service.NewCustomerService(customerRepo)
	customerController := controller.NewCustomerController(customerService)

	// Initialize Order
	orderRepo := repo.NewOrderRepository(mysqlDB)
	orderService := service.NewOrderService(orderRepo)
	orderController := controller.NewOrderController(orderService)

	// Initialize OrderDetail
	orderDetailRepo := repo.NewOrderDetailRepository(mysqlDB)
	orderDetailService := service.NewOrderDetailService(orderDetailRepo)
	orderDetailController := controller.NewOrderDetailController(orderDetailService)

	api := r.Group("/api/v1")
	{
		customers := api.Group("/customer")
		{
			customers.GET("/", customerController.GetAllCustomers)
			customers.GET("/:id", customerController.GetCustomerByID)
			customers.GET("/raw/:id", customerController.GetRawQueryCustomer)
		}

		order := api.Group("/order")
		{
			order.POST("/", orderController.CreateOrder)
			order.GET("/", orderController.GetAllOrders)
			order.GET("/:id", orderController.GetOrderByID)
		}

		orderDetail := api.Group("/order-detail")
		{
			orderDetail.POST("/", orderDetailController.SaveOrderDetail)
			orderDetail.GET("/", orderDetailController.GetAllOrderDetails)
			orderDetail.GET("/:id", orderDetailController.GetOrderDetail)
		}
	}

	return r
}
