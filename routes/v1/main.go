package v1

import "github.com/gofiber/fiber/v2"

func SetupV1Routes(router fiber.Router) {

	SetupOperationRoutes(router)

}
