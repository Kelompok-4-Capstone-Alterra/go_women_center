package helper

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	ValidateStruct(data interface{}) error
}

type playgroundValidator struct {
	validate *validator.Validate
}

func NewPlaygroundValidator(vld *validator.Validate) *playgroundValidator {
	return &playgroundValidator{
		vld,
	}
}

func (v *playgroundValidator) ValidateStruct(data interface{}) error {
	err := v.validate.Struct(data)
	if err != nil {
		return constant.ErrInvalidInput
	}
	return nil
}