package exception

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	HTTPStatusCode int         `json:"-"`
	Code           string      `json:"code"`
	Message        string      `json:"message"`
	Errors         interface{} `json:"error,omitempty"`
	Data           interface{} `json:"data,omitempty"`
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var resp *ErrorResponse
	switch err.(type) {
	case BadRequestError:
		resp = DefaultErrBadRequest
		resp.Errors = err.Error()
	case NotFoundError:
		resp = DefaultErrNotFound
		resp.Errors = err.Error()
	case UnauthorizedError:
		resp = DefaultErrUnauthenticated
		resp.Errors = err.Error()
	default:
		resp = &DefaultErrorResponse
		resp.Errors = err.Error()
	}
	var e *fiber.Error
	if errors.As(err, &e) {
		resp.Code = "500"
		resp.Message = e.Message
	}

	return ctx.Status(resp.HTTPStatusCode).JSON(resp)
}
