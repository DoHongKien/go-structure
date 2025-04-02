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

	limit, err := strconv.Atoi(ctx.Query("size"))
	if err != nil {
		limit = 10 // Default limit
	}

	offset, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		offset = 0 // Default offset
	}

	customers, err := c.service.GetAllCustomers(limit, (offset-1)*limit)
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

func (c *CustomerController) GetRawQueryCustomer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := c.service.GetRawQueryCustomer(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}
