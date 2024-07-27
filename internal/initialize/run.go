package initialize

import (
	"github.com/longtk26/go-ecommerce/global"
	"go.uber.org/zap"
)


func Run() {
	LoadConfig()	
	InitLogger()
	global.Logger.Info("Config loaded successfully", zap.String("config ok", "ok"))
	InitMysql()	
	InitRedis()
	r := InitRouter()

	r.Run(":8080")
}