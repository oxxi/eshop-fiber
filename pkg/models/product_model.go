package models

import "github.com/go-playground/validator"

type ProductModel struct {
	Id          uint    `json:"id,omitempty"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Quantity    uint64  `json:"quantity" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
