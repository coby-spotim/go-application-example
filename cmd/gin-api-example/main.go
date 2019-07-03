package main

import (
	"os"

	"github.com/coby-spotim/gin-example/internal/gin-api-example/handlers"
	"github.com/coby-spotim/gin-example/internal/gin-api-example/registries"
)

func main() {
	registry := registries.NewSimpleRegistry(getEnv())
	router := handlers.SetupRouter(registry)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func getEnv() string {
	env := os.Getenv("GO_ENV")
	if os.Getenv("GO_ENV") == "" {
		env = "development"
	}
	return env
}
