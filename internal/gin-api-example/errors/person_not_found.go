package errors

import "fmt"

type ErrPersonNotFound struct {
	Name string
}

func (e *ErrPersonNotFound) Error() string {
	return fmt.Sprintf("The name %s was not found in the db", e.Name)
}
