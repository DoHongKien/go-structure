package user

import (
	"github.com/DoHongKien/go-structure/internal/controller"
	"github.com/DoHongKien/go-structure/internal/repo"
	"github.com/DoHongKien/go-structure/internal/service"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct{}

func (or *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {

	orderRepository := repo.NewOrderRepository()
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(orderService)

	orderRouter := Router.Group("/order")
	{
		orderRouter.GET("/", orderController.GetAllOrders)
		orderRouter.GET("/:id", orderController.GetOrderByID)
		orderRouter.POST("/", orderController.CreateOrder)
	}
}
