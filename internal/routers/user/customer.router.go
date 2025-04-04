package user

import (
	"github.com/DoHongKien/go-structure/internal/controller"
	"github.com/DoHongKien/go-structure/internal/repo"
	"github.com/DoHongKien/go-structure/internal/service"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (cr *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {

	customerRepository := repo.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)

	customerRouter := Router.Group("/customer")
	{
		customerRouter.GET("/", customerController.GetAllCustomers)
		customerRouter.GET("/:id", customerController.GetCustomerByID)
		customerRouter.GET("/raw/:id", customerController.GetRawQueryCustomer)
	}
}
