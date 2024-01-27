package list

import (
	"log"
	"net/http"
	"todos/ports/gettodos"
	"todos/utils/errs"
)

type service struct {
	getTodos gettodos.Port
}

func New(getTodos gettodos.Port) Service {
	return &service{
		getTodos: getTodos,
	}
}

func (s *service) Execute(request Request) (*Response, errs.AppError) {
	if err := request.Validate(); err != nil {
		log.Println(err)
		return nil, errs.New(http.StatusBadRequest, "E001", "")
	}

	todos, err := s.getTodos.Execute()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "E002", "")
	}

	return NewResponseInfo(*todos).
		Filter(request.Filter).
		Sort(request.Sort), nil
}
