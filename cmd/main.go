package main

import (
	"digital-library-dashboard/config"
	"digital-library-dashboard/pkg/utils"

	"digital-library-dashboard/internal/routes"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pressly/goose"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: could not load .env file: %v", err)
	}

	db := config.ConnectDB()

	defer db.Close()

	if err := goose.Up(db, "./migrations/db"); err != nil {
		log.Fatal(err)
	}

	r := fiber.New()
	r.Use(utils.Logger())

	auth := r.Group("/api")

	routes.SetupRoute(auth, db)

	log.Fatal(r.Listen(":8081"))
}
