package utils

import (
	"github.com/gofiber/fiber/v2"
)

type TokenResponse struct {
	Token string `json:"token" example:"abc123"`
}

type SuccessResponse struct {
	Status  int         `json:"status" example:"200"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status" example:"500"`
	Message string `json:"message" example:"Internal Server Error"`
}

// SuccessResponse sends a success response with the given status, message, and data.
// @Description This function is used to send a successful response in JSON format.
func SuccessResponseFunc(c *fiber.Ctx, status int, message string, data interface{}) error {
	response := SuccessResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return c.Status(status).JSON(response)
}

// ErrorResponse sends an error response with the given status and message.
// @Description This function is used to send an error response in JSON format.
func ErrorResponseFunc(c *fiber.Ctx, status int, message string) error {
	response := ErrorResponse{
		Status:  status,
		Message: message,
	}
	return c.Status(status).JSON(response)
}
