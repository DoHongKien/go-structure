package user

import (
	"github.com/DoHongKien/go-structure/internal/controller"
	"github.com/DoHongKien/go-structure/internal/repo"
	"github.com/DoHongKien/go-structure/internal/service"
	"github.com/gin-gonic/gin"
)

type OrderDetailRouter struct{}

func (od *OrderDetailRouter) InitOrderDetailRouter(Router *gin.RouterGroup) {

	orderDetailRepository := repo.NewOrderDetailRepository()
	orderDetailService := service.NewOrderDetailService(orderDetailRepository)
	orderDetailController := controller.NewOrderDetailController(orderDetailService)

	orderDetailRouter := Router.Group("/order-detail")
	{
		orderDetailRouter.GET("/", orderDetailController.GetAllOrderDetails)
		orderDetailRouter.GET("/:id", orderDetailController.GetOrderDetail)
		orderDetailRouter.POST("/", orderDetailController.SaveOrderDetail)
	}
}
