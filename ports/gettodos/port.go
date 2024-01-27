package gettodos

import "todos/ports/bizmodel"

type Port interface {
	Execute() (*Response, error)
}

type Response = []bizmodel.Todo
