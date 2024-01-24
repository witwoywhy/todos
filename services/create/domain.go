package create

import "todos/utils/errs"

type Service interface {
	Execute(request Request) (*Response, errs.AppError)
}

type Request struct {
	Title       string  `json:"title" validate:"required,max=100"`
	Description *string `json:"description"`
	Image       *string `json:"image" validate:"omitempty,base64"`
	Status      string  `json:"status" validate:"required,status"`
	Date        string  `json:"date" validate:"required,RFC3339"`
}

type Response struct {
	ID string `json:"id"`
}
