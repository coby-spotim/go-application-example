package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/health" {
			ctx.String(http.StatusOK, "OK")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
