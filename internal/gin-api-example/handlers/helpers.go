package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/probably-not/go-application-example/internal/gin-api-example/registries"
)

func withRegistry(handler func(*gin.Context, *registries.SimpleRegistry), registry *registries.SimpleRegistry) func(*gin.Context) {
	return func(ctx *gin.Context) {
		handler(ctx, registry)
	}
}
