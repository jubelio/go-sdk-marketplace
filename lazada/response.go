package lazada

type BaseResponse struct {
	Code       string `json:"code"`
	Success    string `json:"success"`
	ErrCode    string `json:"err_code"`
	RequestID  string `json:"request_id"`
	ErrMessage string `json:"err_message"`
}
