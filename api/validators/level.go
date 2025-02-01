package validators

import (
	"github.com/Nearrivers/dnd-grid-server/api/presenter"
	"github.com/go-playground/validator/v10"
)

var levelValidator *validator.Validate

func ValidateLevel(l presenter.Level) error {
	levelValidator = validator.New(validator.WithRequiredStructEnabled())
	err := levelValidator.Struct(l)
	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	return err.(validator.ValidationErrors)[0]
}
