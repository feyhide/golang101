package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/feyhide/golang101/28_webapp/internal/types"
	"github.com/feyhide/golang101/28_webapp/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// creating a student
		var student types.Student

		err := json.NewDecoder(req.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			// if null body
			response.WriteJson(res, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(res, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(res, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		response.WriteJson(res, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
