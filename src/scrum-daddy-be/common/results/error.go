package results

import (
	"net/http"
	"time"
)

type ErrType string

const (
	ValidationErrType ErrType = "Validation error"
	NotFoundErrType   ErrType = "Entity Not Found"
)

type ErrorResult struct {
	Code      int       `json:"code"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	TimeStamp time.Time `json:"timeStamp"`
	Type      ErrType   `json:"type"`
}

func (e *ErrorResult) Error() string {
	return e.Title
}

func NewErrorResult(code int, title string, detail string) *ErrorResult {
	return &ErrorResult{
		Code:      code,
		Title:     title,
		Detail:    detail,
		TimeStamp: time.Now(),
		Type:      GetDefaultErrType(code),
	}
}

func NewTypedErrorResult(code int, title string, detail string, errType ErrType) *ErrorResult {
	if errType == "" {
		errType = GetDefaultErrType(code)
	}
	return &ErrorResult{
		Code:      code,
		Title:     title,
		Detail:    detail,
		TimeStamp: time.Now(),
		Type:      errType,
	}
}

func GetDefaultErrType(code int) ErrType {
	switch code {
	case http.StatusBadRequest:
		return "Bad Request"
	case http.StatusUnauthorized:
		return "Unauthorized"
	case http.StatusForbidden:
		return "Forbidden"
	case http.StatusNotFound:
		return "Not Found"
	case http.StatusConflict:
		return "Conflict"
	}
	return "Internal Server Error"
}

func ValidationError(title string, detail string, errType ErrType) *ErrorResult {
	return NewTypedErrorResult(400, title, detail, errType)
}

func NotFoundError(title string, detail string, errType ErrType) *ErrorResult {
	return NewTypedErrorResult(400, title, detail, errType)
}
