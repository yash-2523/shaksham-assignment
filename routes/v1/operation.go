package v1

import (
	"github.com/gofiber/fiber/v2"

	v1 "shaksham/controllers/v1"
)

func SetupOperationRoutes(router fiber.Router) {
	operation := router.Group("/operation")

	operation.Post("/", v1.CreateOperation)
}
