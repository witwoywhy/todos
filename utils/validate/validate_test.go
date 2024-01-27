package validate_test

import (
	"testing"
	"todos/utils/constants"
	"todos/utils/validate"

	"github.com/stretchr/testify/assert"
)

func TestValidateStatus(t *testing.T) {
	validate.InitValidate()

	type request struct {
		Status string `validate:"status"`
	}

	type testCase struct {
		name     string
		given    request
		isPassed bool
	}

	testCases := []testCase{
		{
			name: "passed",
			given: request{
				Status: constants.StatusCompleted,
			},
			isPassed: true,
		},
		{
			name: "failed",
			given: request{
				Status: "DONE",
			},
			isPassed: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate.Validator.Struct(tc.given)
			assert.Equal(t, tc.isPassed, err == nil)
		})
	}
}

func TestRFC3339(t *testing.T) {
	validate.InitValidate()

	type request struct {
		Date string `validate:"RFC3339"`
	}

	type testCase struct {
		name     string
		given    request
		isPassed bool
	}

	testCases := []testCase{
		{
			name: "passed",
			given: request{
				Date: "2024-01-08T00:00:00+07:00",
			},
			isPassed: true,
		},
		{
			name: "failed",
			given: request{
				Date: "2024-01-08T00:00:00Z07:00",
			},
			isPassed: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate.Validator.Struct(tc.given)
			assert.Equal(t, tc.isPassed, err == nil)
		})
	}
}

func TestSortField(t *testing.T) {
	validate.InitValidate()

	type request struct {
		Field string `validate:"sortField"`
	}

	type testCase struct {
		name     string
		given    request
		isPassed bool
	}

	testCases := []testCase{
		{
			name: "passed",
			given: request{
				Field: "status",
			},
			isPassed: true,
		},
		{
			name: "failed",
			given: request{
				Field: "image",
			},
			isPassed: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate.Validator.Struct(tc.given)
			assert.Equal(t, tc.isPassed, err == nil)
		})
	}
}

func TestSortOrder(t *testing.T) {
	validate.InitValidate()

	type request struct {
		Order string `validate:"sortOrder"`
	}

	type testCase struct {
		name     string
		given    request
		isPassed bool
	}

	testCases := []testCase{
		{
			name: "passed",
			given: request{
				Order: "desc",
			},
			isPassed: true,
		},
		{
			name: "passed",
			given: request{
				Order: "asc",
			},
			isPassed: true,
		},
		{
			name: "failed",
			given: request{
				Order: "descasc",
			},
			isPassed: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate.Validator.Struct(tc.given)
			assert.Equal(t, tc.isPassed, err == nil)
		})
	}
}
