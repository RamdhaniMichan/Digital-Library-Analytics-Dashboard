package main

import (
	"digital-library-dashboard/config"
	bookHandler "digital-library-dashboard/internal/book/handler"
	bookRepository "digital-library-dashboard/internal/book/repository"
	bookService "digital-library-dashboard/internal/book/service"

	userHandler "digital-library-dashboard/internal/user/handler"
	userRepository "digital-library-dashboard/internal/user/repository"
	userService "digital-library-dashboard/internal/user/service"

	memberHandler "digital-library-dashboard/internal/member/handler"
	memberRepository "digital-library-dashboard/internal/member/repository"
	memberService "digital-library-dashboard/internal/member/service"
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

	// r.Post("/login", userHandler.LoginHandler(db))

	auth := r.Group("/api")

	bookRepository := bookRepository.NewRepository(db)
	bookService := bookService.NewService(bookRepository)
	bookHandler.RegisterRoutes(auth, bookService)

	userRepository := userRepository.NewRepository(db)
	userService := userService.NewService(userRepository)
	userHandler.RegisterRoutes(auth, userService)

	memberRepository := memberRepository.NewRepository(db)
	memberService := memberService.NewService(memberRepository)
	memberHandler.RegisterRoutes(auth, memberService)

	// lending.RegisterRoutes(auth, db)
	// analytics.RegisterRoutes(auth, db)

	log.Fatal(r.Listen(":8081"))
}
