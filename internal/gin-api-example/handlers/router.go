package handlers

import (
	"github.com/coby-spotim/go-application-example/internal/gin-api-example/middlewares"
	"github.com/coby-spotim/go-application-example/internal/gin-api-example/registries"
	"github.com/gin-gonic/gin"
)

func SetupRouter(registry *registries.SimpleRegistry) *gin.Engine {
	router := gin.New()
	router.Use(middlewares.ResponseTime(), middlewares.Health(), middlewares.MatchedPath())

	// Get/Create User
	router.GET("/user/:name", withRegistry(GetPerson, registry))
	router.POST("/user", withRegistry(SetPerson, registry))

	return router
}
