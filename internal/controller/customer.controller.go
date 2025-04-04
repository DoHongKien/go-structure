package controller

import (
	"strconv"

	"github.com/DoHongKien/go-structure/internal/service"
	"github.com/DoHongKien/go-structure/pkg/response"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService service.ICustomerService
}

func NewCustomerController(customerService service.ICustomerService) *CustomerController {
	return &CustomerController{customerService: customerService}
}

func (c *CustomerController) GetAllCustomers(ctx *gin.Context) {

	limit, err := strconv.Atoi(ctx.Query("size"))
	if err != nil {
		limit = 10 // Default limit
	}

	offset, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		offset = 1 // Default offset
	}

	customers, err := c.customerService.GetAllCustomers(limit, (offset-1)*limit)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, customers)
}

func (c *CustomerController) GetCustomerByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}
	customer, err := c.customerService.GetCustomerByID(id)

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, customer)
}

func (c *CustomerController) GetRawQueryCustomer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}

	customer, err := c.customerService.GetRawQueryCustomer(id)

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, customer)
}
