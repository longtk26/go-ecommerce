package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/longtk26/go-ecommerce/global"
	"github.com/longtk26/go-ecommerce/internal/routers"
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

	// Middlewares
	// r.Use()		//Logging
	// r.Use()		//Cross
	// r.Use()     //Limiter global

	manageRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1")
	{
		MainGroup.GET("/checkStatus", func(ctx *gin.Context) {})	//Monitoring
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitUserRouter(MainGroup)
	}
  
	return r
}

