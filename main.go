package main

import (
	"log"
	"shaksham/config"
	"shaksham/helpers"
	"shaksham/middlewares"
	"shaksham/routes"
	"shaksham/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	utils.LoadEnv()

	db, err := config.InitDbConn()
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	helpers.ProcessOldJobs()

	app := fiber.New(fiber.Config{
		AppName:   "Shaksham Assignment",
		BodyLimit: 5 * 1024 * 1024,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))
	app.Use(logger.New())
	app.Use(recover.New())

	app.Use(middlewares.DBMiddleware(db))

	routes.SetupRoutes(app)

	port := utils.GetEnv("PORT")
	if port == "" {
		port = ":8000"
	}
	// Start server
	log.Println("Starting server on :" + port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
