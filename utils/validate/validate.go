package validate

import (
	"fmt"
	"todos/utils/constants"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func InitValidate() {
	Validator = validator.New()

	if err := Validator.RegisterValidation("status", validateStatus); err != nil {
		panic(fmt.Errorf("failed to register custom validate status: %v", err))
	}
}

func validateStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	if status != constants.StatusCompleted && status != constants.StatusInProgress {
		return false
	}

	return true
}
