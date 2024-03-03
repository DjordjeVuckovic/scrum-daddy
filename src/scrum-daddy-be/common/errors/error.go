package errors

import "time"

type ErrorResult struct {
	Code        int       `json:"code"`
	Message     string    `json:"message"`
	Description string    `json:"description"`
	TimeStamp   time.Time `json:"timeStamp"`
}

func (e *ErrorResult) Error() string {
	return e.Message
}

func NewErrorResult(code int, message string, description string) *ErrorResult {
	return &ErrorResult{
		Code:        code,
		Message:     message,
		Description: description,
		TimeStamp:   time.Now(),
	}
}
