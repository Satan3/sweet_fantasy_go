package validation

import "github.com/go-playground/validator"

type Error struct {
	FailedField string
	Tag         string
	Value       string
}

func Validate(model interface{}) []*Error {
	var errors []*Error
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := Error{
				FailedField: err.Field(),
				Tag:         err.Tag(),
			}
			errors = append(errors, &element)
		}
	}
	return errors
}
