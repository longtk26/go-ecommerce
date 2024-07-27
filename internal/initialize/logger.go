package initialize

import (
	"github.com/longtk26/go-ecommerce/global"
	"github.com/longtk26/go-ecommerce/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}