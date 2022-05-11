package middleware

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Failed string
	Tag    string
	Value  interface{}
}

func StructValidator(models interface{}) []*Validator {
	var errors []*Validator
	
	err := validator.New().Struct(models)
	if err != nil {
		for _, i := range err.(validator.ValidationErrors) {
			var e Validator
			e.Failed = i.StructNamespace()
			e.Tag = i.Tag()
			e.Value = i.Param()
			errors = append(errors, &e)
		}
	}

	return errors
}