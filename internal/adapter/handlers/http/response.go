package http

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewResponse(code int, message string, data any) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func handleError(ctx *fiber.Ctx, code int, err error) error {
	rsp := NewResponse(code, err.Error(), nil)
	return ctx.JSON(fiber.Map{
		"code":    rsp.Code,
		"message": rsp.Message,
		"data":    rsp.Data,
	})
}

func handleResponse(ctx *fiber.Ctx, code int, data any, message string) {
	rsp := NewResponse(code, message, data)
	ctx.JSON(fiber.Map{
		"code":    rsp.Code,
		"message": rsp.Message,
		"data":    rsp.Data,
	})
}
