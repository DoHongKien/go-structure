package controller

import (
	"strconv"

	"github.com/DoHongKien/go-structure/internal/model"
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

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the provided details
// @Tags OrderController
// @Accept json
// @Produce json
// @Param order body model.Order true "Order details"
// @Success 200 {object} response.ResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /orders [post]
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order model.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
		return
	}

	orderResponse, err := c.service.CreateOrder(&order)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, orderResponse)
}

func (c *OrderController) GetOrderByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
		return
	}

	order, err := c.service.GetOrderByID(id)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, order)
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	orders, err := c.service.GetAllOrders()
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, orders)
}
