package resterr

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_generic(t *testing.T) {
	causeUsername := Causes{
		Field:   "username",
		Message: "Invalid username",
	}
	causeEmail := Causes{
		Field:   "Email",
		Message: "Invalid email",
	}

	causes := []Causes{}
	causes = append(causes, causeUsername)
	causes = append(causes, causeEmail)

	err := NewError("Testing generate error", "generic_error", http.StatusBadRequest, causes)

	assert.NotNil(t, err, "Expect an error from NewError function")
	assert.Equal(t, err.Message, "Testing generate error", "Expect the message to be the same as stated")
	assert.Equal(t, err.Err, "generic_error", "Expect the message to be the same as stated")
	assert.Equal(t, http.StatusBadRequest, err.Code, "Expect the code to be a bed request error")
	assert.Greater(t, len(err.Causes), 0, "Expected validation error causes")
	assert.Equal(t, len(err.Causes), 2, "Expected two validation error causes")
	assert.Error(t, err, "Expect to be an error")
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("Testing bad request error")

	assert.NotNil(t, err, "Expect error not be nil")
	assert.Equal(t, err.Message, "Testing bad request error", "Expect the message to be the same as stated")
	assert.Equal(t, err.Err, "bad_request", "Expect the error to be the same as stated")
	assert.Equal(t, http.StatusBadRequest, err.Code, "Expect the error code to be a bed request error")
	assert.Error(t, err, "Expect to be an error")
}

func TestNewBadRequestValidationError(t *testing.T) {
	causeUsername := Causes{
		Field:   "username",
		Message: "Invalid username",
	}
	causeEmail := Causes{
		Field:   "Email",
		Message: "Invalid email",
	}

	causes := []Causes{}
	causes = append(causes, causeUsername)
	causes = append(causes, causeEmail)

	err := NewBadRequestValidationError("Testing validation error", causes)

	assert.NotNil(t, err, "Expect error not be nil")
	assert.Equal(t, err.Message, "Testing validation error", "Expect message to be the same what was declared in")
	assert.Equal(t, err.Err, "bad_request", "Expect the error to be the same as stated")
	assert.Equal(t, http.StatusBadRequest, err.Code, "Expect the error code to be a bed request error")
	assert.Greater(t, len(err.Causes), 0, "Expected validation error causes")
	assert.Equal(t, len(err.Causes), 2, "Expected two validation error causes")
	assert.Equal(t, err.Error(), "Testing validation error")
	assert.Error(t, err, "Expect to be an error")
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("Testing internal server error")

	assert.NotNil(t, err, "Expect error not be nil")
	assert.Equal(t, err.Message, "Testing internal server error", "Expect the message to be the same as stated")
	assert.Equal(t, err.Err, "internernal_server_error", "Expect the error to be the same as stated")
	assert.Equal(t, http.StatusInternalServerError, err.Code, "Expect the error code to be an internal server error")
	assert.Error(t, err, "Expect to be an error")
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("Testing not found error")

	assert.NotNil(t, err, "Expect error not be nil")
	assert.Equal(t, err.Message, "Testing not found error", "Expect the message to be the same as stated")
	assert.Equal(t, err.Err, "not_found", "Expect the error to be the same as stated")
	assert.Equal(t, http.StatusNotFound, err.Code, "Expect the error code to be a not found error")
	assert.Error(t, err, "Expect to be an error")
}

func TestNewForbiddenError(t *testing.T) {
	err := NewForbiddenError("Testing forbidden error")

	assert.NotNil(t, err, "Expect error not be nil")
	assert.Equal(t, err.Message, "Testing forbidden error", "Expect the message to be the same as stated")
	assert.Equal(t, err.Err, "forbidden", "Expect the error to be the same as stated")
	assert.Equal(t, http.StatusForbidden, err.Code, "Expect the error code to be a forbidden error")
	assert.Error(t, err, "Expect to be an error")
}
