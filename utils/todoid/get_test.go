package todoid_test

import (
	"testing"
	"todos/utils/todoid"

	"github.com/stretchr/testify/assert"
)

func TestGetID(t *testing.T) {
	id := todoid.GetID()
	assert.NotNil(t, id)
}
