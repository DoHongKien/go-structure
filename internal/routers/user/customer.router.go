package user

import (
	"fmt"

	"github.com/DoHongKien/go-structure/global"
	"github.com/DoHongKien/go-structure/internal/wire"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (cr *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {

	// customerRepository := repo.NewCustomerRepository()
	// customerService := service.NewCustomerService(customerRepository)
	customerController, err := wire.InitCustomerRouterHandler()

	if err != nil {
		global.Logger.Error(fmt.Sprintf("Failed to initialize customer router handler: %v", err))
	}

	customerRouter := Router.Group("/customer")
	{
		customerRouter.GET("/", customerController.GetAllCustomers)
		customerRouter.GET("/:id", customerController.GetCustomerByID)
		customerRouter.GET("/raw/:id", customerController.GetRawQueryCustomer)
	}
}
