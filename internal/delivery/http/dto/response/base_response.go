// 6. BASE RESPONSE
// File: internal/delivery/http/dto/response/base_response.go

package response

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(message string, data interface{}) BaseResponse {
	return BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, err error) BaseResponse {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	return BaseResponse{
		Success: false,
		Message: message,
		Error:   errMsg,
	}
}
