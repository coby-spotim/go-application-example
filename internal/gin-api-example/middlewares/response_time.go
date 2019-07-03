package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func ResponseTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		ctx.Next()
		if matchedPath, ok := ctx.Get("matched_path"); ok {
			log.Println("Response Time For", matchedPath, ": ", time.Since(t).Seconds())
		}
	}
}
