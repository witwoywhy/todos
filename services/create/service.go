package create

import (
	"log"
	"net/http"
	"todos/ports/createtodo"
	"todos/utils/errs"
)

type service struct {
	createTodo createtodo.Port
}

func New(createTodo createtodo.Port) Service {
	return &service{
		createTodo: createTodo,
	}
}

func (s *service) Execute(request Request) (*Response, errs.AppError) {
	if err := request.Validate(); err != nil {
		log.Println(err)
		return nil, errs.New(http.StatusBadRequest, "E001", "")
	}

	createTodoRequest := request.ToCreateTodoRequest()
	err := s.createTodo.Execute(createTodoRequest)
	if err != nil {
		log.Println(err)
		return nil, errs.New(http.StatusInternalServerError, "E002", "")
	}

	return &Response{
		ID: createTodoRequest.ID,
	}, nil
}
