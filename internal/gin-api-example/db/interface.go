package db

import (
	"github.com/coby-spotim/gin-example/pkg/types"
)

type DB interface {
	GetPerson(name string) (types.Person, error)
	SetPerson(types.Person) error
}
