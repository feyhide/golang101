package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

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
			slog.Error("empty request body", slog.String("error", err.Error()))
			response.WriteJson(res, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			slog.Error("failed to decode request body", slog.String("error", err.Error()))
			response.WriteJson(res, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//validation
		if err := validator.New().Struct(student); err != nil {
			slog.Error("validation failed", slog.String("error", err.Error()))
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(res, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		if err != nil {
			slog.Error("failed to create student", slog.String("error", err.Error()))
			response.WriteJson(res, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		slog.Info("user created", slog.String("userId", fmt.Sprint(lastId)))

		response.WriteJson(res, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")
		slog.Info("fetching student", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("invalid student ID", slog.String("error", err.Error()))
			response.WriteJson(res, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		student, err := storage.GetById(intId)
		if err != nil {
			slog.Error("failed to fetch student", slog.String("error", err.Error()))
			response.WriteJson(res, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(res, http.StatusOK, student)
	}
}
