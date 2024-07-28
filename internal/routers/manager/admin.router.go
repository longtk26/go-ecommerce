package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct {}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// Public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login", func(ctx *gin.Context) {})
	}
	
	// Private router
	adminRouterPrivate := Router.Group("/admin/user")
	{
		adminRouterPrivate.POST("/user_info", func(ctx *gin.Context) {})
	}
}