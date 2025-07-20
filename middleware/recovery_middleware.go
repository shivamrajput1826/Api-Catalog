package middleware

import (
	"fmt"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/shivamrajput1826/api-catalog/logger"
)

func RecoveryMiddleware(c *fiber.Ctx) error {
	customLogger := logger.CreateLogger("RecoveryMiddleware").WithFiberContext(c)
	defer func() {
		if r := recover(); r != nil {
			customLogger.Debug("Recovered from panic", map[string]interface{}{
				"message":    fmt.Sprintf("Recovered from panic: %v", r),
				"stackTrace": string(debug.Stack()),
			})
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}
	}()
	return c.Next()
}
