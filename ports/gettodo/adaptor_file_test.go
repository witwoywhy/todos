package gettodo_test

import (
	"os"
	"testing"
	"todos/infra"
	"todos/ports/bizmodel"
	"todos/ports/gettodo"
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
		request gettodo.Request
	}

	type when struct {
		initFileFunc  func() *os.File
		changeFileDir *string
	}

	type expect struct {
		response *gettodo.Response
		isError  bool
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
				request: gettodo.Request{
					ID: "60d809b2-33e4-4439-b789-644e29042304",
				},
			},
			when: &when{
				initFileFunc: initFile,
			},
			expect: &expect{
				response: &bizmodel.Todo{
					ID:                "60d809b2-33e4-4439-b789-644e29042304",
					Title:             "t1",
					Description:       ptr.ToPtr("d1"),
					Image:             ptr.ToPtr("ZG8gdGVzdCBodWdlbWFu"),
					Status:            "IN_PROGRESS",
					Date:              "2024-01-08T00:00:00+07:00",
					CreatedAtDatetime: "2024-01-27T19:16:33+07:00",
				},
			},
		},
		{
			name: "failed when read file",
			given: &given{
				request: gettodo.Request{
					ID: "60d809b2-33e4-4439-b789-644e29042304",
				},
			},
			when: &when{
				initFileFunc:  initFile,
				changeFileDir: ptr.ToPtr(""),
			},
			expect: &expect{
				isError: true,
			},
		},
		{
			name: "failed when byte to object",
			given: &given{
				request: gettodo.Request{
					ID: "60d809b2-33e4-4439-b789-644e29042304",
				},
			},
			when: &when{
				initFileFunc: initFileWrongFormatJson,
			},
			expect: &expect{
				isError: true,
			},
		},
		{
			name: "failed when mapping id",
			given: &given{
				request: gettodo.Request{
					ID: "uuid",
				},
			},
			when: &when{
				initFileFunc: initFile,
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

			getTodo := gettodo.NewAdaptorFile(infra.NewFileManager(fileName))
			response, err := getTodo.Execute(tc.given.request)
			if tc.expect.isError {
				assert.NotNil(t, err)
			}
			assert.Equal(t, tc.expect.response, response)

			os.Remove(file.Name())
		})
	}
}
