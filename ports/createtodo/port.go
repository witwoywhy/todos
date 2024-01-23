package createtodo

import "todos/ports/bizmodel"

type Port interface {
	Execute(request Request) error
}

type Request = bizmodel.Todo
