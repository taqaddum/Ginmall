package response

import (
	"GinMall/serializer/status"
)

var okMsg = status.Msg[status.Ok]
var failedMsg = status.Msg[status.Error]

type BasicResponse struct {
	Code    int    `json:"status,omitempty"`
	Message string `json:"msg,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Success() *BasicResponse {
	return &BasicResponse{
		Code:    status.Ok,
		Message: okMsg,
	}
}

func SuccessMsg(msg string) *BasicResponse {
	return &BasicResponse{
		Code:    status.Ok,
		Message: msg,
	}
}

func SuccessData(data any) *BasicResponse {
	return &BasicResponse{
		Code:    status.Ok,
		Message: okMsg,
		Data:    data,
	}
}

func SuccessMsgAndData(msg string, data any) *BasicResponse {
	return &BasicResponse{
		Code:    status.Ok,
		Message: msg,
		Data:    data,
	}
}

func Failed() *BasicResponse {
	return &BasicResponse{
		Code:    status.Error,
		Message: failedMsg,
	}
}

func FailedCode(code int) *BasicResponse {
	return &BasicResponse{
		Code:    code,
		Message: status.Msg[code],
	}
}

func FailedMsg(msg string) *BasicResponse {
	return &BasicResponse{
		Code:    status.Error,
		Message: msg,
	}
}

func FailedCodeAndMsg(code int, msg string) *BasicResponse {
	return &BasicResponse{
		Code:    code,
		Message: msg,
	}
}
