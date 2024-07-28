package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// Public router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/search", func(ctx *gin.Context) {})
		productRouterPublic.GET("/detail/:id", func(ctx *gin.Context) {})
	}

	// Private router
}