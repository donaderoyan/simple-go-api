package util

import (
	"github.com/gin-gonic/gin"
)

type Responses struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ResponsesPagination struct {
	Pagination interface{} `json:"pagination"`
	Responses
}

type ErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Error      interface{} `json:"error"`
}

func APIResponsePagination(ctx *gin.Context, Message string, StatusCode int, Method string, Data interface{}, Pagination interface{}) {

	jsonResponse := ResponsesPagination{
		Responses: Responses{
			StatusCode: StatusCode,
			Method:     Method,
			Message:    Message,
			Data:       Data,
		},
		Pagination: Pagination,
	}
	if StatusCode >= 400 {
		ctx.JSON(StatusCode, jsonResponse)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Method string, Data interface{}) {

	jsonResponse := Responses{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.JSON(StatusCode, jsonResponse)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func ValidatorErrorResponse(ctx *gin.Context, StatusCode int, Method string, Error interface{}) {
	errResponse := ErrorResponse{
		StatusCode: StatusCode,
		Method:     Method,
		Error:      Error,
	}

	ctx.JSON(StatusCode, errResponse)
	defer ctx.AbortWithStatus(StatusCode)
}
