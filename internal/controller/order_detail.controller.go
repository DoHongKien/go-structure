package controller

import (
	"strconv"

	"github.com/DoHongKien/go-structure/internal/models"
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
	var orderDetail *models.OrderDetail

	if err := ctx.ShouldBindJSON(&orderDetail); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	}

	orderDetailResponse, err := c.orderDetailService.SaveOrderDetail(orderDetail)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, orderDetailResponse)
}

func (c *OrderDetailController) GetOrderDetail(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	}
	orderDetail, err := c.orderDetailService.GetOrderDetail(id)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, orderDetail)
}

func (c *OrderDetailController) GetAllOrderDetails(ctx *gin.Context) {
	orderDetails, err := c.orderDetailService.GetAllOrderDetails()
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, orderDetails)
}
