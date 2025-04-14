package initialize

import (
	"fmt"

	"github.com/DoHongKien/go-structure/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
	LoadConfig()
	m := global.Config.Mysql

	fmt.Println("Loading configuration mysql: ", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Config Log ok!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	return r
}
