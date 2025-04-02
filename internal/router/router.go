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

	// Initialize repositories
	customerRepo := repo.NewCustomerRepository(mysqlDB)

	// Initialize services
	customerService := service.NewCustomerService(customerRepo)

	// Initialize controllers
	customerController := controller.NewCustomerController(customerService)

	api := r.Group("/api/v1")
	{
		customers := api.Group("/customers")
		{
			customers.GET("/", customerController.GetAllCustomers)
			customers.GET("/:id", customerController.GetCustomerByID)
			customers.GET("/raw/:id", customerController.GetRawQueryCustomer)
		}
	}

	return r
}
