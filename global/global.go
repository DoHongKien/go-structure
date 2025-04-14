package global

import (
	"github.com/DoHongKien/go-structure/pkg/logger"
	"github.com/DoHongKien/go-structure/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
