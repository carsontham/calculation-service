package errors

import "time"

type ErrorResponse struct {
	Timestamp string `json:"timestamp"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Status    int    `json:"status,omitempty"`
}

// NewErrorResponse creates a new ErrorResponse with the current timestamp.
func NewErrorResponse(err string, msg string, status int) *ErrorResponse {
	return &ErrorResponse{
		Timestamp: time.Now().Format(time.RFC3339),
		Error:     err,
		Message:   msg,
		Status:    status,
	}
}
