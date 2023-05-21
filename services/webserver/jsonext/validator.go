package jsonext

import "github.com/go-playground/validator/v10"

type InputValidator interface {
	Validate(input any) error
}

type inputValidator struct {
	validator *validator.Validate
}

func NewInputValidator() InputValidator {
	return &inputValidator{
		validator: validator.New(),
	}
}

func (v *inputValidator) Validate(input any) error {
	return v.validator.Struct(input)
}
