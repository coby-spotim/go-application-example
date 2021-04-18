package handlers

import (
	"net/http"

	"github.com/probably-not/go-application-example/internal/gin-api-example/errors"

	"github.com/probably-not/go-application-example/internal/gin-api-example/registries"
	"github.com/probably-not/go-application-example/pkg/types"

	"github.com/gin-gonic/gin"
)

func GetPerson(ctx *gin.Context, registry *registries.SimpleRegistry) {
	name := ctx.Param("name")
	person, err := registry.DB.GetPerson(name)
	if err != nil {
		if err, ok := err.(*errors.ErrPersonNotFound); ok {
			ctx.JSON(http.StatusNotFound, gin.H{"name": name, "error": "The name was not found in the database"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}
	ctx.JSON(http.StatusOK, person)
}

func SetPerson(ctx *gin.Context, registry *registries.SimpleRegistry) {
	var person types.Person

	if ctx.Bind(&person) == nil {
		if err := registry.DB.SetPerson(person); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"status": "ok"})
	}
}
