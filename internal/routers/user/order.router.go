package user

import (
	"fmt"

	"github.com/DoHongKien/go-structure/global"
	"github.com/DoHongKien/go-structure/internal/wire"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct{}

func (or *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {

	// orderRepository := repo.NewOrderRepository()
	// orderService := service.NewOrderService(orderRepository)
	orderController, err := wire.InitOrderRouterHandler()

	if err != nil {
		global.Logger.Error(fmt.Sprintf("Failed to initialize order router handler: %v", err))
	}

	orderRouter := Router.Group("/order")
	{
		orderRouter.GET("/", orderController.GetAllOrders)
		orderRouter.GET("/:id", orderController.GetOrderByID)
		orderRouter.POST("/", orderController.CreateOrder)
	}
}
