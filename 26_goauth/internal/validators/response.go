package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "Ok"
	StatusError = "Error"
)

func ValidateError(errs validator.ValidationErrors) Response {
	var errorMessages []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errorMessages = append(errorMessages, fmt.Sprintf("Field %s is required field", err.Field()))
		default:
			errorMessages = append(errorMessages, fmt.Sprintf("Field %s is invalid", err.Field()))
		}
	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errorMessages, ","),
	}
}
