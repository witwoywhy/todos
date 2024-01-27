package validate

import (
	"fmt"
	"time"
	"todos/utils/constants"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func InitValidate() {
	Validator = validator.New()

	if err := Validator.RegisterValidation("status", validateStatus); err != nil {
		panic(fmt.Errorf("failed to register custom validate status: %v", err))
	}

	if err := Validator.RegisterValidation("RFC3339", rfc3339); err != nil {
		panic(fmt.Errorf("failed to register custom validate RFC3339: %v", err))
	}

	if err := Validator.RegisterValidation("sortField", sortField); err != nil {
		panic(fmt.Errorf("failed to register custom validate sortField: %v", err))
	}

	if err := Validator.RegisterValidation("sortOrder", sortOrder); err != nil {
		panic(fmt.Errorf("failed to register custom validate sortOrder: %v", err))
	}
}

func validateStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	if status != constants.StatusCompleted && status != constants.StatusInProgress {
		return false
	}

	return true
}

func rfc3339(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())
	return err == nil
}

func sortField(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	if field != constants.SortTitle && field != constants.SortStatus && field != constants.SortDate {
		return false
	}

	return true
}

func sortOrder(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	if field != constants.SortDesc && field != constants.SortAsc {
		return false
	}

	return true
}
