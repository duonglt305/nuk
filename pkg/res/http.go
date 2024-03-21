package res

import (
	v "duonglt.net/pkg/validator"
	"encoding/json"
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func Error(w http.ResponseWriter, err error) {
	status := http.StatusBadRequest
	resp := ErrorResponse{
		Message: "Oops! Something went wrong. Please try again later.",
		Errors:  make(map[string]string),
	}
	switch {
	case errors.As(err, &v.ValidationError{}):
		status = http.StatusUnprocessableEntity
		resp.Message = err.Error()
		resp.Errors = err.(v.ValidationError).Errors()
		break
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Ok(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		Error(w, err)
	}
}
