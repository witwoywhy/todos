package update

import "todos/utils/errs"

type Service interface {
	Execute(request Request) errs.AppError
}

type Request struct {
	ID          string  `json:"id" validate:"required"`
	Title       string  `json:"title" validate:"required,max=100"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	Status      string  `json:"status" validate:"required,status"`
	Date        string  `json:"date" validate:"required,RFC3339"`
}
