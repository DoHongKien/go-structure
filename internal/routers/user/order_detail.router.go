package user

import (
	"fmt"

	"github.com/DoHongKien/go-structure/global"
	"github.com/DoHongKien/go-structure/internal/wire"
	"github.com/gin-gonic/gin"
)

type OrderDetailRouter struct{}

func (od *OrderDetailRouter) InitOrderDetailRouter(Router *gin.RouterGroup) {

	// orderDetailRepository := repo.NewOrderDetailRepository()
	// orderDetailService := service.NewOrderDetailService(orderDetailRepository)
	orderDetailController, err := wire.InitOrderDetailRouterHandler()

	if err != nil {
		global.Logger.Error(fmt.Sprintf("Failed to initialize order detail router handler: %v", err))
	}

	orderDetailRouter := Router.Group("/order-detail")
	{
		orderDetailRouter.GET("/", orderDetailController.GetAllOrderDetails)
		orderDetailRouter.GET("/:id", orderDetailController.GetOrderDetail)
		orderDetailRouter.POST("/", orderDetailController.SaveOrderDetail)
	}
}
