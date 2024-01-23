package update

import (
	"log"
	"net/http"
	"todos/ports/gettodo"
	"todos/ports/updatetodo"
	"todos/utils/errs"
)

type service struct {
	getTodo    gettodo.Port
	updateTodo updatetodo.Port
}

func New(
	getTodo gettodo.Port,
	updateTodo updatetodo.Port,
) Service {
	return &service{
		getTodo:    getTodo,
		updateTodo: updateTodo,
	}
}

func (s *service) Execute(request Request) errs.AppError {
	if err := request.Validate(); err != nil {
		log.Println(err)
		return errs.New(http.StatusBadRequest, "E001", "")
	}

	todo, err := s.getTodo.Execute(request.ToGetTodoRequest())
	switch err {
	case nil:
		goto UPDATE
	case gettodo.ErrNotFound:
		return errs.New(http.StatusBadRequest, "E001", "")
	default:
		return errs.New(http.StatusInternalServerError, "E002", "")
	}

UPDATE:
	err = s.updateTodo.Execute(buildUpdateTodoRequest(request, todo))
	if err != nil {
		return errs.New(http.StatusInternalServerError, "E002", "")
	}

	return nil
}
