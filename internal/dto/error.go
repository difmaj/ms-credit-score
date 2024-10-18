package dto

import (
	"encoding/json"
	"fmt"
)

// APIError is the API error.
type APIError struct {
	Status  int    `json:"status"`          // code is the API code
	Err     any    `json:"error,omitempty"` // Err is the error that caused the API error
	Message string `json:"message"`         // Message is the API message
}

// MarshalJSON marshals the API error.
func (e APIError) MarshalJSON() ([]byte, error) {
	type Alias APIError
	var errString string
	if e.Err != nil {
		switch v := e.Err.(type) {
		case string:
			errString = v
		case error:
			errString = v.Error()
		default:
			errBytes, err := json.Marshal(v)
			if err != nil {
				errString = fmt.Sprintf("unknown error: %v", v)
			} else {
				errString = string(errBytes)
			}
		}
	}

	return json.Marshal(&struct {
		Err string `json:"error,omitempty"`
		Alias
	}{
		Err:   errString,
		Alias: (Alias)(e),
	})
}

func (e *APIError) Error() string {
	if e.Message == "" {
		return fmt.Sprintf("status %v", e.Err)
	}
	return e.Message
}

// NewAPIError returns a new API error.
func NewAPIError(code int, err error, message string, args ...any) *APIError {
	return &APIError{
		Status:  code,
		Err:     err,
		Message: message,
	}
}

// APIErrors is a list of API errors.
type APIErrors []*APIError

// Error returns the error message.
func (e APIErrors) Error() string {
	if len(e) == 0 {
		return ""
	}
	return e[0].Error()
}
