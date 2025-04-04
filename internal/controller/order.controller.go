package controller

import (
	"strconv"

	"github.com/DoHongKien/go-structure/internal/models"
	"github.com/DoHongKien/go-structure/internal/service"
	"github.com/DoHongKien/go-structure/pkg/response"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service service.IOrderService
}

func NewOrderController(service service.IOrderService) *OrderController {
	return &OrderController{
		service: service,
	}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}

	orderResponse, err := c.service.CreateOrder(&order)

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, orderResponse)
}

func (c *OrderController) GetOrderByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}
	order, err := c.service.GetOrderByID(id)

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, order)
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	orders, err := c.service.GetAllOrders()

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, orders)
}
