package lazada

type BaseResponse struct {
	Code       string `json:"code"`
	Success    bool   `json:"success,omitempty"`
	ErrCode    string `json:"err_code,omitempty"`
	RequestID  string `json:"request_id,omitempty"`
	ErrMessage string `json:"err_message,omitempty"`
}
