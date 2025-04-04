package initialize

import (
	"github.com/DoHongKien/go-structure/global"
	"github.com/DoHongKien/go-structure/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
