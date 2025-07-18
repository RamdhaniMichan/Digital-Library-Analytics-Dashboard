package main

import (
	"digital-library-dashboard/config"
	"digital-library-dashboard/pkg/utils"
	"os"

	_ "digital-library-dashboard/docs"
	"digital-library-dashboard/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"github.com/pressly/goose"
)

// @title Digital Library API
// @version 1.0
// @description This is a digital library server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email your-email@domain.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /api
// @schemes http
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: could not load .env file: %v", err)
	}

	db := config.ConnectDB()

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing DB: %v", err)
		}
	}()

	if err := goose.Up(db, "./migrations/db"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	r := fiber.New()
	r.Use(utils.Logger())
	r.Get("/swagger/*", swagger.HandlerDefault)

	auth := r.Group("/api")

	routes.SetupRoute(auth, db)

	r.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "path not found",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Server is running at http://localhost:%s", port)
	log.Fatal(r.Listen(":" + port))
}
