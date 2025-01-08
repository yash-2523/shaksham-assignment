package routes

import (
	v1 "shaksham/routes/v1"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")

	v1.SetupV1Routes(apiV1)
}
