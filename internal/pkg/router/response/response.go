package response

import (
	"encoding/json"
	"net/http"

	"github.com/difmaj/ms-credit-score/internal/dto"
)

// Response is a generic response struct.
type Response[T any] struct {
	Success bool            `json:"success"`
	Return  T               `json:"return"`
	Errors  []*dto.APIError `json:"errors"`
}

// Empty is an empty struct.
type Empty struct{}

func writeResponse[T any](w http.ResponseWriter, code int, v Response[T]) {
	body, err := json.Marshal(v)
	if err != nil {
		Error(w, &dto.APIError{
			Status:  http.StatusInternalServerError,
			Message: "Failed to marshal response",
			Err:     err,
		})
		return
	}
	w.WriteHeader(code)
	w.Write(body)
}

// Write writes a response.
func Write[T any](w http.ResponseWriter, code int, v Response[T]) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	writeResponse(w, code, v)
}

// Error writes an error response.
func Error(w http.ResponseWriter, errs ...*dto.APIError) {
	res := Response[any]{
		Success: false,
		Errors:  errs,
		Return:  nil,
	}
	Write(w, errs[0].Status, res)
}

// Ok writes a successful response.
func Ok[T any](w http.ResponseWriter, code int, v T) {
	res := Response[T]{
		Success: true,
		Return:  v,
	}
	Write(w, code, res)
}
