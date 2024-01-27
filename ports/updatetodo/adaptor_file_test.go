package updatetodo_test

import (
	"os"
	"testing"
	"todos/infra"
	"todos/ports/bizmodel"
	"todos/ports/updatetodo"
	"todos/utils/ptr"

	"github.com/stretchr/testify/assert"
)

func initFile() *os.File {
	f, err := os.CreateTemp("./", "")
	if err != nil {
		panic(err)
	}

	f.Write([]byte(`{"60d809b2-33e4-4439-b789-644e29042304":{"id":"60d809b2-33e4-4439-b789-644e29042304","title":"t1","description":"d1","image":"ZG8gdGVzdCBodWdlbWFu","status":"IN_PROGRESS","date":"2024-01-08T00:00:00+07:00","createdAtDatetime":"2024-01-27T19:16:33+07:00"}}`))

	return f
}

func initFileWrongFormatJson() *os.File {
	f, err := os.CreateTemp("./", "")
	if err != nil {
		panic(err)
	}

	f.Write([]byte(`{[]}`))

	return f
}

func TestExecute(t *testing.T) {
	type given struct {
		request updatetodo.Request
	}

	type when struct {
		initFileFunc  func() *os.File
		changeFileDir *string
	}

	type expect struct {
		isError bool
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
				request: bizmodel.Todo{
					ID: "60d809b2-33e4-4439-b789-644e29042304",
				},
			},
			when: &when{
				initFileFunc: initFile,
			},
			expect: &expect{
				isError: false,
			},
		},
		{
			name:  "failed when read file",
			given: &given{},
			when: &when{
				initFileFunc:  initFile,
				changeFileDir: ptr.ToPtr(""),
			},
			expect: &expect{
				isError: true,
			},
		},
		{
			name:  "failed when byte to object",
			given: &given{},
			when: &when{
				initFileFunc: initFileWrongFormatJson,
			},
			expect: &expect{
				isError: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file := tc.when.initFileFunc()
			defer file.Close()

			fileName := file.Name()
			if tc.when.changeFileDir != nil {
				fileName = ptr.StringNotNil(tc.when.changeFileDir)
			}

			updateTodo := updatetodo.NewAdaptorFile(infra.NewFileManager(fileName))
			err := updateTodo.Execute(tc.given.request)
			if tc.expect.isError {
				assert.NotNil(t, err)
			}

			os.Remove(file.Name())
		})
	}
}
