package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func MatchedPath() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		matchedPath := ctx.Request.URL.String()
		for _, p := range ctx.Params {
			matchedPath = strings.Replace(matchedPath, p.Value, ":"+p.Key, 1)
		}
		ctx.Set("matched_path", matchedPath)
		ctx.Next()
	}
}
