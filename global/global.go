package global

import (
	"github.com/DoHongKien/go-structure/pkg/logger"
	"github.com/DoHongKien/go-structure/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
