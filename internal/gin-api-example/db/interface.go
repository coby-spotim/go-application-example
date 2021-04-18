package db

import (
	"github.com/probably-not/go-application-example/pkg/types"
)

type DB interface {
	GetPerson(name string) (types.Person, error)
	SetPerson(types.Person) error
}
