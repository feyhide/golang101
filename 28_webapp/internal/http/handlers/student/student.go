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
			response.WriteJson(res, http.StatusBadRequest, false, "Empty request body", nil)
			return
		}

		if err != nil {
			slog.Error("failed to decode request body", slog.String("error", err.Error()))
			response.WriteJson(res, http.StatusBadRequest, false, "Invalid JSON format", nil)
			return
		}

		//validation
		if err := validator.New().Struct(student); err != nil {
			slog.Error("validation failed", slog.String("error", err.Error()))
			response.ValidationError(res, err.(validator.ValidationErrors))
			return
		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		if err != nil {
			slog.Error("failed to create student", slog.String("error", err.Error()))
			response.GeneralError(res, http.StatusInternalServerError, err)
			return
		}
		slog.Info("user created", slog.String("userId", fmt.Sprint(lastId)))

		response.WriteJson(res, http.StatusCreated, true, "Student created successfully", map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")
		slog.Info("fetching student", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("invalid student ID", slog.String("error", err.Error()))
			response.GeneralError(res, http.StatusBadRequest, err)
			return
		}

		student, err := storage.GetById(intId)
		if err != nil {
			slog.Error("failed to fetch student", slog.String("error", err.Error()))
			response.GeneralError(res, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(res, http.StatusOK, true, "Student fetched successfully", student)
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		slog.Info("Getting all students")

		students, err := storage.GetList()
		if err != nil {
			slog.Error("failed to fetch students", slog.String("error", err.Error()))
			response.GeneralError(res, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(res, http.StatusOK, true, "Students fetched successfully", students)
	}
}

func Update(storage storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// updating student
		var updateStudent types.UpdateStudent

		err := json.NewDecoder(req.Body).Decode(&updateStudent)
		if errors.Is(err, io.EOF) {
			// if null body
			slog.Error("empty request body", slog.String("error", err.Error()))
			response.WriteJson(res, http.StatusBadRequest, false, "Empty request body", nil)
			return
		}

		if err != nil {
			slog.Error("failed to decode request body", slog.String("error", err.Error()))
			response.WriteJson(res, http.StatusBadRequest, false, "Invalid JSON format", nil)
			return
		}

		//validation
		if err := validator.New().Struct(updateStudent); err != nil {
			slog.Error("validation failed", slog.String("error", err.Error()))
			response.ValidationError(res, err.(validator.ValidationErrors))
			return
		}

		updatedStudent, err := storage.UpdateStudent(
			updateStudent.Id,
			updateStudent.Name,
			updateStudent.Age,
			updateStudent.Email,
		)

		if err != nil {
			slog.Error("failed to update student", slog.String("error", err.Error()))
			response.GeneralError(res, http.StatusInternalServerError, err)
			return
		}
		response.WriteJson(res, http.StatusOK, true, "Student updated successfully", updatedStudent)
	}
}

func Delete(storage storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// deleting student
		id := req.PathValue("id")
		slog.Info("fetching student", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("invalid student ID", slog.String("error", err.Error()))
			response.GeneralError(res, http.StatusBadRequest, err)
			return
		}

		student, err := storage.DeleteStudent(intId)
		if err != nil {
			slog.Error("failed to delete student", slog.String("error", err.Error()))
			response.GeneralError(res, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(res, http.StatusOK, true, "Student deleted successfully", student)
	}
}
