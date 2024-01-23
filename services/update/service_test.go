package update_test

import (
	"errors"
	"net/http"
	"testing"
	"todos/ports/bizmodel"
	"todos/ports/gettodo"
	"todos/ports/updatetodo"
	"todos/services/update"
	"todos/utils/constants"
	"todos/utils/errs"
	"todos/utils/ptr"
	"todos/utils/validate"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExecute(t *testing.T) {
	type given struct {
		request update.Request
	}

	type when struct {
		getTodoResponse *gettodo.Response
		getTodoError    error

		updateTodoError error
	}

	type expect struct {
		err errs.AppError
	}

	type testCase struct {
		name string

		given  *given
		when   *when
		expect *expect
	}

	testCases := []testCase{
		{
			name: "success",
			given: &given{
				request: update.Request{
					ID:          "6d38f1cf-5e40-42b6-86c6-b05996a963cd",
					Title:       "Update Task",
					Description: ptr.ToPtr("test update"),
					Image:       nil,
					Status:      constants.StatusCompleted,
				},
			},
			when: &when{
				getTodoResponse: &bizmodel.Todo{
					ID: "6d38f1cf-5e40-42b6-86c6-b05996a963cd",
				},
			},
			expect: &expect{
				err: nil,
			},
		},
		{
			name: "invalid request",
			given: &given{
				request: update.Request{},
			},
			when: &when{},
			expect: &expect{
				err: errs.New(http.StatusBadRequest, "E001", ""),
			},
		},
		{
			name: "todo not found",
			given: &given{
				request: update.Request{
					ID:          "6d38f1cf-5e40-42b6-86c6-b05996a963cd",
					Title:       "Update Task",
					Description: ptr.ToPtr("test update"),
					Image:       nil,
					Status:      constants.StatusCompleted,
				},
			},
			when: &when{
				getTodoError: gettodo.ErrNotFound,
			},
			expect: &expect{
				err: errs.New(http.StatusBadRequest, "E001", ""),
			},
		},
		{
			name: "get todo response error",
			given: &given{
				request: update.Request{
					ID:          "6d38f1cf-5e40-42b6-86c6-b05996a963cd",
					Title:       "Update Task",
					Description: ptr.ToPtr("test update"),
					Image:       nil,
					Status:      constants.StatusCompleted,
				},
			},
			when: &when{
				getTodoError: errors.New("test error"),
			},
			expect: &expect{
				err: errs.New(http.StatusInternalServerError, "E002", ""),
			},
		},
		{
			name: "update todo response error",
			given: &given{
				request: update.Request{
					ID:          "6d38f1cf-5e40-42b6-86c6-b05996a963cd",
					Title:       "Update Task",
					Description: ptr.ToPtr("test update"),
					Image:       nil,
					Status:      constants.StatusCompleted,
				},
			},
			when: &when{
				getTodoResponse: &bizmodel.Todo{
					ID: "6d38f1cf-5e40-42b6-86c6-b05996a963cd",
				},
				updateTodoError: errors.New("test error"),
			},
			expect: &expect{
				err: errs.New(http.StatusInternalServerError, "E002", ""),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validate.InitValidate()

			getTodo := gettodo.NewMock()
			getTodo.On("Execute", mock.Anything).Return(tc.when.getTodoResponse, tc.when.getTodoError)

			updateTodo := updatetodo.NewMock()
			updateTodo.On("Execute", mock.Anything).Return(tc.when.updateTodoError)

			svc := update.New(getTodo, updateTodo)
			err := svc.Execute(tc.given.request)
			assert.Equal(t, tc.expect.err, err)
		})
	}
}
