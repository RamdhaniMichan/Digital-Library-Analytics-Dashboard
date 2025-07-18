package main

import (
	"digital-library-dashboard/config"
	"digital-library-dashboard/internal/book/handler"
	"digital-library-dashboard/internal/book/repository"
	"digital-library-dashboard/internal/book/service"
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

	// r.Post("/login", user.LoginHandler(db))

	auth := r.Group("/api")

	bookRepository := repository.NewRepository(db)
	bookService := service.NewService(bookRepository)
	handler.RegisterRoutes(auth, bookService)
	// lending.RegisterRoutes(auth, db)
	// analytics.RegisterRoutes(auth, db)

	log.Fatal(r.Listen(":8081"))
}
