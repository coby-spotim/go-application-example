package db

import (
	"github.com/coby-spotim/go-application-example/pkg/types"
)

type DB interface {
	GetPerson(name string) (types.Person, error)
	SetPerson(types.Person) error
}
