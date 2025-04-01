package controller

import (
	"net/http"
	"strconv"

	"github.com/DoHongKien/go-structure/internal/service"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	service *service.CustomerService
}

func NewCustomerController(service *service.CustomerService) *CustomerController {
	return &CustomerController{service: service}
}

func (c *CustomerController) GetAllCustomers(ctx *gin.Context) {
	customers, err := c.service.GetAllCustomers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customers)
}

func (c *CustomerController) GetCustomerByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	customer, err := c.service.GetCustomerByID(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}
