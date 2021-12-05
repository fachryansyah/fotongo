package baseCommands

import (
	"fotongo/app/utils/dtos"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

// Generate bcrypt hash
func GenerateHash(plain string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func ValidateRequest(source interface{}) *dtos.JSONResponses {
	var errors []*dtos.ErrorResponse
	validate := validator.New()
	err := validate.Struct(source)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element dtos.ErrorResponse
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	response := BadRequestResponse(errors)

	return &response
}
