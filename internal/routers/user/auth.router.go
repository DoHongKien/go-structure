package user

import (
	"github.com/DoHongKien/go-structure/internal/wire"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {}

func (ar *AuthRouter) InitAuthRouter(router *gin.RouterGroup) {

	authController, err := wire.InitAuthRouterHandler()

	if err != nil {
		panic(err)
	}

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/login/gen-token", authController.LoginGenToken)
	}
}