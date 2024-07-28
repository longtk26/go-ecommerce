package global

import (
	"github.com/longtk26/go-ecommerce/pkg/logger"
	"github.com/longtk26/go-ecommerce/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb	*gorm.DB
	Rdb *redis.Client
)