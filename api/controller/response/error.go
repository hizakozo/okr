package response

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func AlreadyExistsData() *ErrorResponse{
	return &ErrorResponse{Code: 401, Message: "already exists data"}
}

func NotFoundData() *ErrorResponse {
	return &ErrorResponse{Code: 402, Message: "not found data"}
}

func FormBindFailed() *ErrorResponse{
	return &ErrorResponse{Code: 403, Message: "failed form bind"}
}

func ValidateFailed() *ErrorResponse {
	return &ErrorResponse{Code: 404, Message: "failed validate"}
}

func OtherError(err error) *ErrorResponse {
	return &ErrorResponse{Code: 411, Message: err.Error()}
}