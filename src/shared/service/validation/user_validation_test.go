package validation

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUserError_UnmarshalTypeError(t *testing.T) {
	jsonError := &json.UnmarshalTypeError{
		Field: "username",
		Type:  nil,
	}

	err := ValidateUserError(jsonError)
	assert.NotNil(t, err, "Expect a error before to call the ValidateUserError function")
	assert.Equal(t, "Invalid field type", err.Message, "Expect message to be 'Invalid field type'")
	assert.Equal(t, http.StatusBadRequest, err.Code, "Expect the error code to be a bad request error")
}

func TestValidateUserError_ValidationErrors(t *testing.T) {
	err := Validate.Struct(struct {
		Username string `validate:"required"`
	}{})

	validationErr := ValidateUserError(err)
	assert.NotNil(t, validationErr, "Expect a error before to call the ValidateUserError function")
	assert.Equal(t, "Some fields are invalid", validationErr.Message, "Expect message to be 'Some fields are invalid'")
	assert.Equal(t, http.StatusBadRequest, validationErr.Code, "Expect the error code to be a bad request error")
	assert.Greater(t, len(validationErr.Causes), 0, "Expected validation error causes")
}

func TestValidateUserError_GenericError(t *testing.T) {
	genericError := errors.New("generic error")

	err := ValidateUserError(genericError)
	assert.NotNil(t, err, "Expect a error before to call the ValidateUserError function")
	assert.Equal(t, "Error trying to convert fields", err.Message, "Expect message to be 'Error trying to convert fields'")
	assert.Equal(t, http.StatusBadRequest, err.Code, "Expect the error code to be a bad request error")
}
