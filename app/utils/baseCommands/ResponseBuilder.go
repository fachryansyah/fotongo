package baseCommands

import (
	"fotongo/app/utils/dtos"
)

func SuccessResponses(data interface{}) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Status:  "Success",
		Message: "Request successfully!",
		Code:    200,
		Data:    data,
	}
}

func NotFoundResponses(message string, data interface{}) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Status:  "Not Found",
		Message: message,
		Code:    404,
		Data:    data,
	}
}

func InternalServerErrorResponses() (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Status:  "Internal Server Error",
		Code:    500,
		Message: "something went wrong!",
		Data:    nil,
	}
}

func ForbiddenResponse(message string) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Status:  "Access Forbidden",
		Message: message,
		Code:    403,
		Data:    nil,
	}
}

func BadRequestResponse(data interface{}) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Status:  "Bad Request",
		Message: "Invalid request body",
		Code:    400,
		Data:    data,
	}
}
