package exception

import "github.com/gofiber/fiber/v2"

func (err *ErrorResponse) Error() interface{} {
	return err.Message
}

var (
	DefaultErrorResponse = ErrorResponse{
		HTTPStatusCode: fiber.StatusInternalServerError,
		Code:           "500",
		Message:        "Internal server error",
	}

	DefaultErrInternalServer = &ErrorResponse{
		HTTPStatusCode: fiber.StatusInternalServerError,
		Code:           "500",
		Message:        "Internal server error",
	}

	DefaultErrBadRequest = &ErrorResponse{
		HTTPStatusCode: fiber.StatusBadRequest,
		Code:           "400",
		Message:        "Bad request",
	}

	DefaultErrPermissionDenied = &ErrorResponse{
		HTTPStatusCode: fiber.StatusForbidden,
		Code:           "403",
		Message:        "Permission denied",
	}

	DefaultErrNotFound = &ErrorResponse{
		HTTPStatusCode: fiber.StatusNotFound,
		Code:           "404",
		Message:        "Not found",
	}

	DefaultErrUnauthenticated = &ErrorResponse{
		HTTPStatusCode: fiber.StatusUnauthorized,
		Code:           "401",
		Message:        "Unauthorized",
	}
)

func PanicLogging(err interface{}) {
	if err != nil {
		panic(err)
	}
}
