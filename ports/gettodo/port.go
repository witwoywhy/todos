package gettodo

import "todos/ports/bizmodel"

type Port interface {
	Execute(request Request) (*Response, error)
}

type Request struct {
	ID string
}

type Response = bizmodel.Todo
