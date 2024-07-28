package user

import "github.com/gin-gonic/gin"

type UserRouter struct {}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", func(ctx *gin.Context) {})
		userRouterPublic.POST("/otp", func(ctx *gin.Context) {})
	}

	// Private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/get_info", func(ctx *gin.Context) {})
	}
}