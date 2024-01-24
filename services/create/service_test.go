package create_test

import (
	"errors"
	"net/http"
	"testing"
	"todos/ports/createtodo"
	"todos/services/create"
	"todos/utils/constants"
	"todos/utils/errs"
	"todos/utils/ptr"
	"todos/utils/todoid"
	"todos/utils/validate"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestExecute(t *testing.T) {
	type given struct {
		request create.Request
	}
	type when struct {
		getIDFunc todoid.GetIDFunc

		createTodoErr error
	}
	type expect struct {
		response *create.Response
		err      errs.AppError
	}

	type testCase struct {
		name   string
		given  *given
		when   *when
		expect *expect
	}

	testCases := []testCase{
		{
			name: "success",
			given: &given{
				request: create.Request{
					Title:       "Do test hugeman",
					Description: ptr.ToPtr("within 7 day"),
					Image:       ptr.ToPtr("ZG8gdGVzdCBodWdlbWFu"),
					Status:      constants.StatusInProgress,
					Date:        "2024-01-24T20:23:39+07:00",
				},
			},
			when: &when{
				getIDFunc: func() string {
					return "9e71128b-d265-4927-ac66-4d7dc32a3514"
				},
				createTodoErr: nil,
			},
			expect: &expect{
				response: &create.Response{
					ID: "9e71128b-d265-4927-ac66-4d7dc32a3514",
				},
			},
		},
		{
			name: "invalid request",
			given: &given{
				request: create.Request{
					Status: "DONE",
				},
			},
			when: &when{},
			expect: &expect{
				err: errs.New(http.StatusBadRequest, "E001", ""),
			},
		},
		{
			name: "craete todo response error",
			given: &given{
				request: create.Request{
					Title:       "Do test hugeman",
					Description: ptr.ToPtr("within 7 day"),
					Image:       ptr.ToPtr("ZG8gdGVzdCBodWdlbWFu"),
					Status:      constants.StatusCompleted,
					Date:        "2024-01-24T20:23:39+07:00",
				},
			},
			when: &when{
				createTodoErr: errors.New("create error"),
			},
			expect: &expect{
				err: errs.New(http.StatusInternalServerError, "E002", ""),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.when.getIDFunc != nil {
				todoid.GetID = tc.when.getIDFunc
			}

			validate.InitValidate()

			createTodo := createtodo.NewMock()
			createTodo.On("Execute", mock.Anything).Return(tc.when.createTodoErr)

			service := create.New(createTodo)

			response, err := service.Execute(tc.given.request)
			if tc.expect.err != nil {
				assert.Equal(t, tc.expect.err, err)
			}

			assert.Equal(t, tc.expect.response, response)
		})
	}
}
