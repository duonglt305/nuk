package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const (
	unknown      = "unknown"
	notfound     = "not_found"
	unauthorized = "unauthorized"
	invalid      = "invalid"
)

func Error(w http.ResponseWriter, err error) {
	switch {
	case errors.As(err, &ValidationError{}):
		UnprocessableEntity(w, err.(ValidationError))
		return
	default:
		BadRequest(w, err)
	}
}

func BadRequest(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	body := map[string]string{
		"message": "Oops! Something went wrong.",
		"code":    unknown,
		"error":   err.Error(),
	}
	log.Printf("error: %+v\n", err)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		Error(w, err)
	}
}

func NotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	body := map[string]string{"message": "Resource not found", "code": notfound}
	if err := json.NewEncoder(w).Encode(body); err != nil {
		Error(w, err)
	}
}

func Unauthorized(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	body := map[string]string{"message": err.Error(), "code": unauthorized}
	if err := json.NewEncoder(w).Encode(body); err != nil {
		Error(w, err)
	}
}

func UnprocessableEntity(w http.ResponseWriter, err ValidationError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	body := map[string]string{"message": err.Error(), "code": invalid}
	if err := json.NewEncoder(w).Encode(body); err != nil {
		Error(w, err)
	}

}

func Ok(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		Error(w, err)
	}
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}
