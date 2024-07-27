package global

import (
	"github.com/longtk26/go-ecommerce/pkg/logger"
	"github.com/longtk26/go-ecommerce/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb	*gorm.DB
)