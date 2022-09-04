package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Controll-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Controll-Allow-Credential", "true")
		ctx.Writer.Header().Set("Access-Controll-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		ctx.Writer.Header().Set("Access-Controll-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
