package initialize

import (
	"fmt"

	"github.com/longtk26/go-ecommerce/global"
)

func Run() {
	// 1. Load config
	LoadConfig()	
	fmt.Println("Load config mysql", global.Config.Mysql.Username)
	InitLogger()
	InitMysql()	
	InitRedis()
	r := InitRouter()

	r.Run(":8080")
}