package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/probably-not/go-application-example/internal/gin-api-example/middlewares"
	"github.com/probably-not/go-application-example/internal/gin-api-example/registries"
)

func SetupRouter(registry *registries.SimpleRegistry) *gin.Engine {
	router := gin.New()
	router.Use(middlewares.ResponseTime(), middlewares.Health(), middlewares.MatchedPath())

	// Get/Create User
	router.GET("/user/:name", withRegistry(GetPerson, registry))
	router.POST("/user", withRegistry(SetPerson, registry))

	return router
}
