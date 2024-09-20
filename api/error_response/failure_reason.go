package error_response

import (
	"encoding/json"
)

// collect here the different error_response
const (
	ErrorUnknownFailureReason            = "UNKNOWN_FAILURE_REASON"                        // error
	PreviewFailureReasonAfterHourInvalid = "PREVIEW_FUTURES_AFTER_HOUR_INVALID_ORDER_TYPE" // preview_failure_reason
)

type ErrorResponse struct {
	Err                  string `json:"error"`
	Message              string `json:"message"`
	ErrorDetails         string `json:"error_details"`
	PreviewFailureReason string `json:"preview_failure_reason"`
}

type ErrorResult struct {
	Success       bool           `json:"success"`
	ErrorResponse *ErrorResponse `json:"error_response"`
}

func Parse(b []byte) *ErrorResponse {
	out := &ErrorResult{}
	err := json.Unmarshal(b, &out)
	if err != nil || out.Success {
		return nil
	}

	return out.ErrorResponse
}

func (e *ErrorResponse) Error() string {
	if e == nil {
		return ""
	}

	return e.Err + " / " + e.PreviewFailureReason
}
