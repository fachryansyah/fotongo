package utils

import (
	"github.com/gofiber/fiber/v2"
)

// ResponseSuccess : returning json structur for success request
func ResponseSuccess(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": message,
		"data":    data,
	})
}

// ResponseNotFound : returning json structur for notfound request
func ResponseNotFound(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(fiber.Map{
		"status":  404,
		"message": message,
	})
}

// ResponseError : returning json structur for error request
func ResponseError(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(fiber.Map{
		"status":  500,
		"message": message,
		"data":    data,
	})
}

// ResponseUnauthenticated : returning json structur for validation error request
func ResponseUnauthenticated(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(fiber.Map{
		"status":  403,
		"message": message,
		"data":    data,
	})
}

// ResponseValidationError : returning json structur for validation error request
func ResponseValidationError(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(fiber.Map{
		"status":  304,
		"message": message,
		"data":    data,
	})
}
