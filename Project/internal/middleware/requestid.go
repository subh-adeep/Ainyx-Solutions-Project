package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Get("X-Request-Id")
		if id == "" {
			id = uuid.New().String()
		}
		c.Set("X-Request-Id", id)
		c.Locals("requestid", id)
		return c.Next()
	}
}
