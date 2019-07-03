package db

import (
	"github.com/coby-spotim/gin-example/internal/gin-api-example/errors"
	"github.com/coby-spotim/gin-example/pkg/types"
)

type Mock struct {
	data map[string]types.Person
}

func NewMock() *Mock {
	return &Mock{
		data: make(map[string]types.Person),
	}
}

func (db *Mock) GetPerson(name string) (types.Person, error) {
	person, ok := db.data[name]
	if ok {
		return person, nil
	} else {
		return person, &errors.ErrPersonNotFound{Name: name}
	}
}

func (db *Mock) SetPerson(person types.Person) error {
	db.data[person.Name] = person
	return nil
}
