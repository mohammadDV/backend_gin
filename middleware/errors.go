package middleware

import "github.com/gin-gonic/gin"

func NotMethodHandler() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "Method not allowed"})
	}

}

func NotRouteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "Route not defined"})
	}
}
