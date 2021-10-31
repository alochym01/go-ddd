package errs

import "net/http"

type AppErr struct {
	Code    int    `json:-`
	Message string `json:"message"`
}

func NotFound(msg string) *AppErr {
	return &AppErr{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func ServerErr(msg string) *AppErr {
	return &AppErr{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}
