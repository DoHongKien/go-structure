package controller

import (
	"github.com/DoHongKien/go-structure/internal/model/dto"
	"github.com/DoHongKien/go-structure/internal/service"
	"github.com/DoHongKien/go-structure/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(service service.IAuthService) *AuthController {
	return &AuthController{
		authService: service,
	}
}

func (ac *AuthController) Login(ctx *gin.Context) {

	var userRequest dto.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}

	result, err := ac.authService.Login(userRequest)

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}

	if result {
		response.SuccessResponse(ctx, response.ErrCodeSuccess, "Login successfully!!!")
	} else {
		response.ErrorResponse(ctx, response.ErrCodeFailed, "Username or password invalid")
	}
}

func (ac *AuthController) LoginGenToken(ctx *gin.Context) {
	var userRequest dto.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}

	result, err := ac.authService.LoginGenToken(userRequest)

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeFailed, err.Error())
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}
