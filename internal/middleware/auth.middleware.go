package middleware

import (
	"github.com/BaoTo12/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valid-token" {
			response.FailureResponse(ctx, 30001)
			ctx.Abort()
		}
		ctx.Next()
	}
}
