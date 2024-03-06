package helper

import (
	"github.com/revandpratama/go-simple-blog-api/dto"
)

type ResponseWithData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(param dto.ResponseParam) any {
	var status string
	var response any

	if param.StatusCode >= 200 && param.StatusCode < 400 {
		status = "success"
	} else {
		status = "error"
	}

	if param.Data != nil {
		response = ResponseWithData{
			Code:    param.StatusCode,
			Status:  status,
			Message: param.Message,
			Data:    param.Data,
		}
	} else {
		response = ResponseWithoutData{
			Code:    param.StatusCode,
			Status:  status,
			Message: param.Message,
		}
	}

	return response
}
