package list

import (
	"todos/ports/bizmodel"
	"todos/utils/errs"
)

type Service interface {
	Execute(request Request) (*Response, errs.AppError)
}

type Request struct {
	Filter Filter `json:"filter"`
	Sort   Sort   `json:"sort"`
}

type Filter struct {
	Title       *string `json:"title" validate:"omitempty,max=100"`
	Description *string `json:"description"`
}

type Sort struct {
	Field *string `json:"field" validate:"omitempty,sortField"`
	Order *string `json:"order" validate:"omitempty,sortOrder"`
}

type Response = []bizmodel.Todo
