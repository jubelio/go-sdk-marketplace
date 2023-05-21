package shopee

type BaseResponse struct {
	RequestID string `json:"request_id"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Warning   string `json:"warning"`
}
