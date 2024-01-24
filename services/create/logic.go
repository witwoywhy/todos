package create

import (
	"time"
	"todos/ports/createtodo"
	"todos/utils/todoid"
	"todos/utils/validate"
)

func (r *Request) Validate() error {
	if err := validate.Validator.Struct(r); err != nil {
		return err
	}

	return nil
}

func (r *Request) ToCreateTodoRequest() createtodo.Request {
	return createtodo.Request{
		ID:                todoid.GetID(),
		Title:             r.Title,
		Description:       r.Description,
		Image:             r.Image,
		Status:            r.Status,
		Date:              r.Date,
		CreatedAtDatetime: time.Now().Format(time.RFC3339),
	}
}
