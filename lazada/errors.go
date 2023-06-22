package lazada

import "encoding/json"

// Error response is used to return as much data as possible to the calling application to help with dealing with any API issues.
type ResponseError struct {
	Code      string          `json:"code"`
	Type      string          `json:"type"`
	Message   string          `json:"message"`
	RequestID string          `json:"request_id"`
	Detail    []*ErrorDetails `json:"detail,omitempty"`
}

type ErrorDetails struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ResponseError) Error() string {
	err := struct {
		Code      string          `json:"code"`
		Type      string          `json:"type"`
		Message   string          `json:"message"`
		RequestID string          `json:"request_id"`
		Detail    []*ErrorDetails `json:"detail"`
	}{
		Code:      e.Code,
		Type:      e.Type,
		Message:   e.Message,
		RequestID: e.RequestID,
		Detail:    e.Detail,
	}

	jsonErr, _ := json.Marshal(err)
	return string(jsonErr)
}
