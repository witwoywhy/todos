package list

import (
	"sort"
	"time"
	"todos/ports/bizmodel"
	"todos/ports/gettodos"
	"todos/utils/constants"
	"todos/utils/ptr"
	"todos/utils/sorted"
	"todos/utils/validate"
)

func (r *Request) Validate() error {
	if err := validate.Validator.Struct(r); err != nil {
		return err
	}

	return nil
}

type ResponseInfo struct {
	todos []bizmodel.Todo
}

func NewResponseInfo(todos gettodos.Response) *ResponseInfo {
	return &ResponseInfo{
		todos: todos,
	}
}

func (r *ResponseInfo) Filter(filter Filter) *ResponseInfo {
	if filter.Title != nil || filter.Description != nil {
		var todos []bizmodel.Todo = make([]bizmodel.Todo, 0)
		for _, todo := range r.todos {
			if todo.Title == ptr.StringNotNil(filter.Title) || ptr.StringNotNil(todo.Description) == ptr.StringNotNil(filter.Description) {
				todos = append(todos, todo)
				continue
			}
		}

		r.todos = todos
	}

	return r
}

func (r *ResponseInfo) Sort(s Sort) *Response {
	switch ptr.StringNotNil(s.Order) {
	case sorted.Asc:
		sorted.String = sorted.StringAsc
		sorted.Date = sorted.DateAsc
	case sorted.Desc:
		sorted.String = sorted.StringDesc
		sorted.Date = sorted.DateDesc
	}

	field := ptr.StringNotNil(s.Field)
	if field == constants.SortTitle {
		sort.Slice(r.todos, func(i, j int) bool {
			return sorted.String(r.todos[i].Title, r.todos[j].Title)
		})
	}

	if field == constants.SortStatus {
		sort.Slice(r.todos, func(i, j int) bool {
			return sorted.String(r.todos[i].Status, r.todos[j].Status)
		})
	}

	if field == constants.SortDate {
		sort.Slice(r.todos, func(i, j int) bool {
			dateI, _ := time.Parse(time.RFC3339, r.todos[i].Date)
			dateJ, _ := time.Parse(time.RFC3339, r.todos[j].Date)
			return sorted.Date(dateI, dateJ)
		})
	}

	return &r.todos
}
