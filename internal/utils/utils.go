package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParseUintID(idStr string) (uint, error) {
	if idStr == "" {
		return 0, fiber.NewError(fiber.StatusBadRequest, "ID parameter is required")
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Invalid ID format")
	}

	return uint(id), nil
}
