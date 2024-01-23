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
