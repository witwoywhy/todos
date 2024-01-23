package update

import (
	"time"
	"todos/ports/bizmodel"
	"todos/ports/gettodo"
	"todos/ports/updatetodo"
	"todos/utils/validate"
)

func (r *Request) Validate() error {
	if err := validate.Validator.Struct(r); err != nil {
		return err
	}

	return nil
}

func (r *Request) ToGetTodoRequest() gettodo.Request {
	return gettodo.Request{
		ID: r.ID,
	}
}

func buildUpdateTodoRequest(request Request, todo *bizmodel.Todo) updatetodo.Request {
	todo.Title = request.Title
	todo.Description = request.Description
	todo.Image = request.Image
	todo.Status = request.Status
	todo.CreatedAtDatetime = time.Now().Format(time.RFC3339)
	return *todo
}
