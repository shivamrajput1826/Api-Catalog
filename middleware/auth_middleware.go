package middleware

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/shivamrajput1826/api-catalog/common"
	"github.com/shivamrajput1826/api-catalog/config"
	"github.com/shivamrajput1826/api-catalog/logger"
)

func AuthMiddleware(c *fiber.Ctx) error {
	customLogger := logger.CreateLogger("AuthMiddleware").WithFiberContext(c)
	authHeader := c.Get("Authorization")
	clientId := c.Get("client-id")
	secret := config.GetConfigValue("JWT_SECRET")
	if secret == "" || authHeader == "" || clientId == "" {
		customLogger.Debug("Missing required authentication parameters")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		customLogger.Debug("Invalid token", "error", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || (!slices.Contains(common.ClientIdsConfig, clientId)) || claims["user_id"] == "" {
		customLogger.Debug("Invalid token", "error", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	c.Locals("user_id", claims["user_id"])
	c.Locals("email", claims["email"])
	c.Locals("role", claims["role"])
	return c.Next()
}
