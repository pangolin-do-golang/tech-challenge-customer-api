package errutil

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	e := &Error{Message: "test message"}
	assert.Equal(t, "test message", e.Error())
}

func TestNewBusinessError(t *testing.T) {
	origErr := errors.New("original error")
	e := NewBusinessError(origErr, "business error occurred")

	assert.Equal(t, "business error occurred", e.Message)
	assert.Equal(t, "BUSINESS", e.Type)
	assert.Equal(t, origErr, e.originalError)
}

func TestNewInputError(t *testing.T) {
	origErr := errors.New("input error")
	e := NewInputError(origErr)

	assert.Equal(t, "INPUT", e.Type)
	assert.Equal(t, origErr, e.originalError)
}
