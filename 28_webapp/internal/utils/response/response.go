package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func WriteJson(res http.ResponseWriter, status int, data any) error {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	err := json.NewEncoder(res).Encode(data)
	if err != nil {
		GeneralError(err)
	}
	return nil
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
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

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
