package handlers

import (
	"github.com/coby-spotim/gin-example/internal/gin-api-example/registries"
	"github.com/gin-gonic/gin"
)

func withRegistry(handler func(*gin.Context, *registries.SimpleRegistry), registry *registries.SimpleRegistry) func(*gin.Context) {
	return func(ctx *gin.Context) {
		handler(ctx, registry)
	}
}
