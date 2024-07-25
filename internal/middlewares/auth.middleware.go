package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/longtk26/go-ecommerce/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token != "valid-token" {
			response.ErrorResponse(
				ctx,
				response.ErrInvalidToken,
			)
			ctx.Abort()
			return
		} 
		ctx.Next()
	}
}