package dto

type ResponseMeta struct {
	Success      bool   `json:"success"`
	MessageTitle string `json:"messageTitle"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
}

type ErrorResponse struct {
	ResponseMeta
	Data   any `json:"data"`
	Errors any `json:"errors,omitempty"`
}

func DefaultErrorResponse() ErrorResponse {
	return DefaultErrorResponseWithMessage("")
}

func DefaultErrorResponseWithMessage(msg string) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success:      false,
			MessageTitle: "Oops, something went wrong.",
			Message:      msg,
			ResponseTime: "",
		},
		Data: nil,
	}
}

func DefaultBadRequestResponse() ErrorResponse {
	return DefaultErrorResponseWithMessage("Bad request")
}
