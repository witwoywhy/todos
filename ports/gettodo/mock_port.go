package gettodo

import (
	"github.com/stretchr/testify/mock"
)

type adaptorMock struct {
	mock.Mock
}

func NewMock() *adaptorMock {
	return &adaptorMock{}
}

func (a *adaptorMock) Execute(request Request) (*Response, error) {
	args := a.Called(request)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Response), nil
}
