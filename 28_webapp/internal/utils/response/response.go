package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func WriteJson(res http.ResponseWriter, status int, success bool, message string, data any) error {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)

	return json.NewEncoder(res).Encode(Response{
		Success: success,
		Message: message,
		Data:    data,
	})
}

func GeneralError(res http.ResponseWriter, status int, err error) error {
	return WriteJson(res, status, false, err.Error(), nil)
}

func ValidationError(res http.ResponseWriter, errs validator.ValidationErrors) error {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("Field %s is required", err.Field()))
		case "email":
			errMsgs = append(errMsgs, fmt.Sprintf("Field %s must be a valid email address", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("Field %s is invalid", err.Field()))
		}
	}

	return WriteJson(res, http.StatusBadRequest, false, "Validation failed", errMsgs)
}
