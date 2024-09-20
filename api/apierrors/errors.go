package apierrors

import (
	"github.com/pkg/errors"
)

type ApiError struct {
	Error        string `json:"error"`
	ErrorDetails string `json:"error_details"`
	Message      string `json:"message"`
}

var (
	NotFoundError         = errors.Errorf("Not found")
	UnauthorizedError     = errors.Errorf("Unauthorized")
	MethodNotAllowedError = errors.Errorf("Method Not Allowed")
)

type ApiError2 struct {
	Error   string `json:"error"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Details struct {
		TypeUrl string `json:"type_url"`
		Value   string `json:"value"`
	} `json:"details"`
}
