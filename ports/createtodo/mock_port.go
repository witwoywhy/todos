package createtodo

import (
	"todos/ports/bizmodel"

	"github.com/stretchr/testify/mock"
)

type adaptorMock struct {
	mock.Mock
}

func NewMock() *adaptorMock {
	return &adaptorMock{}
}

func (a *adaptorMock) Execute(request bizmodel.Todo) error {
	args := a.Called(request)
	if args.Error(0) != nil {
		return args.Error(0)
	}

	return nil
}
