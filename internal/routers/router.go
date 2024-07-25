package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	c "github.com/longtk26/go-ecommerce/internal/controllers"
	"github.com/longtk26/go-ecommerce/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		fmt.Println("Before --> AA")
		ctx.Next()
		fmt.Println("After --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		fmt.Println("Before --> BB")
		ctx.Next()
		fmt.Println("After --> BB")
	}
}

func CC(ctx *gin.Context) {
	fmt.Println("Before --> CC")
	ctx.Next()
	fmt.Println("After --> CC")
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	// Use middleware
	r.Use(middlewares.AuthenMiddleware(), BB(), CC)

	v1 := r.Group("/v1")
	{
	  v1.GET("/ping", c.NewUserController().GetUserByID)
	}
  
	return r
}

