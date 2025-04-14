package initialize

import (
	"github.com/DoHongKien/go-structure/global"
	"github.com/DoHongKien/go-structure/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/api/v2")
	{
		userRouter.InitAuthRouter(MainGroup)         // Initialize the authentication router
		userRouter.InitCustomerRouter(MainGroup)    // Initialize the customer router
		userRouter.InitOrderRouter(MainGroup)       // Initialize the order router
		userRouter.InitOrderDetailRouter(MainGroup) // Initialize the order detail router
	}

	return r
}
