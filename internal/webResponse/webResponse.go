package webResponse

type WebResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewErrorResponse(code int, message string) *WebResponse {
	return &WebResponse{
		Code:    code,
		Message: message,
	}
}

func NewSuccessResponse(code int, message string, data interface{}) *WebResponse {
	return &WebResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
