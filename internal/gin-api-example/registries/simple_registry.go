package registries

import (
	"github.com/coby-spotim/go-application-example/internal/gin-api-example/db"
)

type SimpleRegistry struct {
	Env string
	DB  db.DB
}

func NewSimpleRegistry(env string) *SimpleRegistry {
	return &SimpleRegistry{
		Env: env,
		DB:  db.NewMock(),
	}
}
