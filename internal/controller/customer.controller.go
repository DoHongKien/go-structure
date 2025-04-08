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

// GetAllCustomers godoc
// @Summary Get all customers
// @Description Retrieve a paginated list of all customers
// @Tags CustomerController
// @Accept json
// @Produce json
// @Param size query int false "Number of items per page" default(10)
// @Param page query int false "Page number" default(1)
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /customers [get]
func (c *CustomerController) GetAllCustomers(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("size"))
	if err != nil || limit <= 0 {
		limit = 10 // Default limit
	}

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || page <= 0 {
		page = 1 // Default page
	}

	offset := (page - 1) * limit

	customers, err := c.customerService.GetAllCustomers(limit, offset)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, customers)
}

// GetCustomerByID godoc
// @Summary Get customer by ID
// @Description Retrieve a customer by their ID
// @Tags CustomerController
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /customers/{id} [get]
func (c *CustomerController) GetCustomerByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, "Invalid customer ID")
		return
	}

	customer, err := c.customerService.GetCustomerByID(id)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, customer)
}

// GetRawQueryCustomer godoc
// @Summary Get customer by ID using raw query
// @Description Retrieve a customer by their ID using a raw SQL query
// @Tags CustomerController
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /customers/raw/{id} [get]
func (c *CustomerController) GetRawQueryCustomer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, "Invalid customer ID")
		return
	}

	customer, err := c.customerService.GetRawQueryCustomer(id)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, customer)
}
