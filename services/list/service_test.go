package list_test

import (
	"errors"
	"net/http"
	"testing"
	"todos/ports/bizmodel"
	"todos/ports/gettodos"
	"todos/services/list"
	"todos/utils/errs"
	"todos/utils/ptr"
	"todos/utils/validate"

	"github.com/go-playground/assert/v2"
)

var todo1 = bizmodel.Todo{
	ID:                "7180b8d5-a65f-4d44-a5d3-d25b983856ee",
	Title:             "t1",
	Description:       ptr.ToPtr("d1"),
	Status:            "IN_PROGRESS",
	Date:              "2024-01-01T00:00:00+07:00",
	CreatedAtDatetime: "2024-01-24T22:06:58+07:00",
}
var todo2 = bizmodel.Todo{
	ID:                "89433d45-c76a-4689-8608-85c210b63dc5",
	Title:             "t2",
	Description:       ptr.ToPtr("d2"),
	Status:            "IN_PROGRESS",
	Date:              "2024-01-02T00:00:00+07:00",
	CreatedAtDatetime: "2024-01-24T22:06:50+07:00",
}
var todo3 = bizmodel.Todo{
	ID:                "abd4b523-aae4-45c2-8157-8bc464c68fc4",
	Title:             "t3",
	Description:       ptr.ToPtr("d3"),
	Status:            "COMPLETED",
	Date:              "2024-01-03T00:00:00+07:00",
	CreatedAtDatetime: "2024-01-24T22:06:18+07:00",
}

var getTodosResponse = &[]bizmodel.Todo{todo1, todo2, todo3}

func TestExecute(t *testing.T) {
	type given struct {
		request list.Request
	}

	type when struct {
		getTodosResponse *gettodos.Response
		getTodosError    error
	}

	type expect struct {
		response *list.Response
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
			name: "success with filter title and description & sort date desc",
			given: &given{
				request: list.Request{
					Filter: list.Filter{
						Title:       ptr.ToPtr("t1"),
						Description: ptr.ToPtr("d3"),
					},
					Sort: list.Sort{
						Field: ptr.ToPtr("date"),
						Order: ptr.ToPtr("desc"),
					},
				},
			},
			when: &when{
				getTodosResponse: getTodosResponse,
			},
			expect: &expect{
				response: &[]bizmodel.Todo{todo3, todo1},
			},
		},
		{
			name: "success without filter & sort title asc",
			given: &given{
				request: list.Request{
					Filter: list.Filter{},
					Sort: list.Sort{
						Field: ptr.ToPtr("title"),
						Order: ptr.ToPtr("asc"),
					},
				},
			},
			when: &when{
				getTodosResponse: getTodosResponse,
			},
			expect: &expect{
				response: &[]bizmodel.Todo{todo1, todo2, todo3},
			},
		},
		{
			name: "success without filter & sort status asc",
			given: &given{
				request: list.Request{
					Filter: list.Filter{},
					Sort: list.Sort{
						Field: ptr.ToPtr("status"),
						Order: ptr.ToPtr("asc"),
					},
				},
			},
			when: &when{
				getTodosResponse: &[]bizmodel.Todo{todo2, todo3},
			},
			expect: &expect{
				response: &[]bizmodel.Todo{todo3, todo2},
			},
		},
		{
			name: "invalid request",
			given: &given{
				request: list.Request{
					Filter: list.Filter{
						Title: ptr.ToPtr("aaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbbc"),
					},
					Sort: list.Sort{},
				},
			},
			when: &when{},
			expect: &expect{
				err: errs.New(http.StatusBadRequest, "E001", ""),
			},
		},
		{
			name: "getTodos response error",
			given: &given{
				request: list.Request{
					Filter: list.Filter{},
					Sort:   list.Sort{},
				},
			},
			when: &when{
				getTodosError: errors.New("test error"),
			},
			expect: &expect{
				err: errs.New(http.StatusInternalServerError, "E002", ""),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validate.InitValidate()

			getTodos := gettodos.NewMock()
			getTodos.On("Execute").Return(tc.when.getTodosResponse, tc.when.getTodosError)

			svc := list.New(getTodos)

			response, err := svc.Execute(tc.given.request)
			if tc.expect.err != nil {
				assert.Equal(t, tc.expect.err, err)
			}

			assert.Equal(t, tc.expect.response, response)
		})
	}
}
