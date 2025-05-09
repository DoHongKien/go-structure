package controller

import (
	"strconv"

	"github.com/DoHongKien/go-structure/internal/model"
	"github.com/DoHongKien/go-structure/internal/service"
	"github.com/DoHongKien/go-structure/pkg/response"
	"github.com/gin-gonic/gin"
)

type OrderDetailController struct {
	orderDetailService service.IOrderDetailService
}

func NewOrderDetailController(orderDetailService service.IOrderDetailService) *OrderDetailController {
	return &OrderDetailController{
		orderDetailService: orderDetailService,
	}
}

func (c *OrderDetailController) SaveOrderDetail(ctx *gin.Context) {
	var orderDetail model.OrderDetail

	if err := ctx.ShouldBindJSON(&orderDetail); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
		return
	}

	orderDetailResponse, err := c.orderDetailService.SaveOrderDetail(&orderDetail)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, orderDetailResponse)
}

func (c *OrderDetailController) GetOrderDetail(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, "Invalid ID format")
		return
	}

	orderDetail, err := c.orderDetailService.GetOrderDetail(id)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, orderDetail)
}

func (c *OrderDetailController) GetAllOrderDetails(ctx *gin.Context) {
	orderDetails, err := c.orderDetailService.GetAllOrderDetails()
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, orderDetails)
}
