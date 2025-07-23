package response

import (
	"os"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/helper"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status   string      `json:"status"  example:"success"`
	Code     int         `json:"code" example:"0"`
	Message  string      `json:"message" example:"Success"`
	Data     interface{} `json:"data" `
	MetaData interface{} `json:"metadata"`
}

type MetaData struct {
	Timestamp string `json:"timestamp" example:"2006-01-02T15:04:05Z07:00"`
	Version   string `json:"version" example:"v1"`
	Token     string `json:"token,omitempty"`
}

type MetaDataPaging struct {
	Timestamp   string `json:"timestamp"`
	Version     string `json:"version"`
	Page        int    `json:"page"`
	CurrentPage int64  `json:"current_page"`
	TotalItems  int64  `json:"total_items"`
}

type MetaDataFullPaging struct {
	Timestamp   string `json:"timestamp"`
	Version     string `json:"version"`
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	CurrentPage int64  `json:"current_page"`
	TotalItems  int64  `json:"total_items"`
}

func ResponseSuccess(data interface{}, message string) Response {
	return Response{
		Status:  "success'",
		Code:    0,
		Message: message,
		Data:    data,
		MetaData: MetaData{
			Timestamp: helper.TimeGenerator(),
			Version:   "v1",
		},
	}

}

func ResponseSuccessWithToken(data interface{}, message string, token string) Response {
	return Response{
		Status:  "success'",
		Code:    0,
		Message: message,
		Data:    data,
		MetaData: MetaData{
			Timestamp: helper.TimeGenerator(),
			Version:   "v1",
			Token:     token,
		},
	}

}

func ResponseSuccessWithPaging(data interface{}, message string, page int, currentPage int64, total int64) Response {
	return Response{
		Status:  "success'",
		Code:    0,
		Message: message,
		Data:    data,
		MetaData: MetaDataPaging{
			Timestamp:   helper.TimeGenerator(),
			Version:     "v1",
			Page:        page,
			CurrentPage: currentPage,
			TotalItems:  total,
		},
	}

}

func ResponseSuccessWithFullPaging(data interface{}, message string, page int, currentPage int64, total int64, limit int) Response {
	return Response{
		Status:  "success'",
		Code:    0,
		Message: message,
		Data:    data,
		MetaData: MetaDataFullPaging{
			Timestamp:   helper.TimeGenerator(),
			Version:     "v1",
			Page:        page,
			Limit:       limit,
			CurrentPage: currentPage,
			TotalItems:  total,
		},
	}

}

func WriteResponse(ctx *fiber.Ctx, resp Response, err domain.ErrorData, statusCode int) {
	if err.Code != 0 {
		if os.Getenv("DEBUG") != "dev" {
			err.Errors = nil

		}

		ctx.Status(statusCode).JSON(err)

	} else {
		ctx.Status(statusCode).JSON(resp)
	}

}
