package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/feyhide/golang101/28_webapp/internal/storage"
	"github.com/feyhide/golang101/28_webapp/internal/types"
	"github.com/feyhide/golang101/28_webapp/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
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

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("user created", slog.String("userId", fmt.Sprint(lastId)))

		if err != nil {
			response.WriteJson(res, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(res, http.StatusCreated, map[string]int64{"id": lastId})
	}
}
