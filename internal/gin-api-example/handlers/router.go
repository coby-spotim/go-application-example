package handlers

import (
	"github.com/coby-spotim/gin-example/internal/gin-api-example/registries"
	"github.com/gin-gonic/gin"
)

func SetupRouter(registry *registries.SimpleRegistry) *gin.Engine {
	router := gin.Default()

	// Simple Ping/Pong
	router.GET("/ping", Ping)

	// Get/Create User
	router.GET("/user/:name", withRegistry(GetPerson, registry))
	router.POST("/user", withRegistry(SetPerson, registry))

	return router
}
