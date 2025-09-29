package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/suraj-9849/Golang.git/internal/storage"
	"github.com/suraj-9849/Golang.git/internal/types"
	"github.com/suraj-9849/Golang.git/internal/utils/response"
)

// New returns an HTTP handler for creating a new student.
// It decodes the request body into a Student struct, validates the input,
// and returns appropriate JSON responses.
func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		// Decode the JSON request body into the student struct
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			// Handle empty request body
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err != nil {
			// Handle invalid JSON
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Validate the student struct using go-playground/validator
		if err := validator.New().Struct(student); err != nil {
			// If validation fails, return validation errors
			validatorErrors := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validatorErrors))
			return
		}

		// Log the creation of the student
		slog.Info("Creating the Student")
		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		slog.Info("User Created Successfully!", slog.String("userId", fmt.Sprint(lastId)))
		// Respond with success
		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": (lastId)})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract "id" from query parameters
		idStr := r.PathValue("id")
		if idStr == "" {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("missing id parameter")))
			return
		}

		// Convert idStr to int64
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("invalid id parameter")))
			return
		}

		student, err := storage.GetStudentById(id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		student, err := storage.GetList()
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusOK, student)
	}
}

func DeleteTheStudentByID(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		if idStr == "" {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("missing id parameter")))
			return
		}

		// Convert idStr to int64
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("invalid id parameter")))
			return
		}
		student, err := storage.DeleteTheStudentByID(id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}

func UpdateById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		if idStr == "" {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("missing id parameter")))
			return
		}

		// Convert idStr to int64
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("invalid id parameter")))
			return
		}

		// Parse request body for updated student data
		var studentData types.Student
		err = json.NewDecoder(r.Body).Decode(&studentData)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(errors.New("empty request body")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Validate the student struct
		if err := validator.New().Struct(studentData); err != nil {
			validatorErrors := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validatorErrors))
			return
		}

		student, err := storage.UpdateById(id, studentData.Name, studentData.Email, studentData.Age)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}
