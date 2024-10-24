package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	resterr "github.com/lukinhas563/gochat/src/shared/service/restErr"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)

		transl, _ = unt.GetTranslator("en")

		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationError error) *resterr.RestError {
	var jsonError *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonError) {
		return resterr.NewBadRequestError("Invalid field type")
	} else if errors.As(validationError, &jsonValidationError) {
		errorCauses := []resterr.Causes{}

		for _, e := range validationError.(validator.ValidationErrors) {
			cause := resterr.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return resterr.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return resterr.NewBadRequestError("Error trying to convert fields")
	}
}
