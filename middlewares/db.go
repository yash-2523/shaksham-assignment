package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DBMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Store the GORM db instance in the context
		c.Locals("db", db)

		return c.Next()
	}
}
