package manager

import "github.com/gin-gonic/gin"

type UserRouter struct {}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Private router
	userRouterPrivate := Router.Group("/admin/user")
	{
		userRouterPrivate.POST("/active_user", func(ctx *gin.Context) {})
	}
}