package routes

import (
	"database/sql"

	bookHandler "digital-library-dashboard/internal/book/handler"
	bookRepository "digital-library-dashboard/internal/book/repository"
	bookService "digital-library-dashboard/internal/book/service"
	"digital-library-dashboard/internal/middleware"

	"github.com/gofiber/fiber/v2"

	userHandler "digital-library-dashboard/internal/user/handler"
	userRepository "digital-library-dashboard/internal/user/repository"
	userService "digital-library-dashboard/internal/user/service"

	memberHandler "digital-library-dashboard/internal/member/handler"
	memberRepository "digital-library-dashboard/internal/member/repository"
	memberService "digital-library-dashboard/internal/member/service"

	lendingHandler "digital-library-dashboard/internal/lending/handler"
	lendingRepository "digital-library-dashboard/internal/lending/repository"
	lendingService "digital-library-dashboard/internal/lending/service"
)

func SetupRoute(auth fiber.Router, db *sql.DB) {
	userRepository := userRepository.NewRepository(db)
	userService := userService.NewService(userRepository)
	userHandler.RegisterRoutes(auth.Group("/v1"), userService)

	v1 := auth.Group("/v1", middleware.JWTMiddleware())

	bookRepository := bookRepository.NewRepository(db)
	bookService := bookService.NewService(bookRepository)
	bookHandler.RegisterRoutes(v1, bookService)

	memberRepository := memberRepository.NewRepository(db)
	memberService := memberService.NewService(memberRepository)
	memberHandler.RegisterRoutes(v1, memberService)

	lendingRepository := lendingRepository.NewLendingRepository(db)
	lendingService := lendingService.NewLendingService(lendingRepository, bookRepository, memberRepository)
	lendingHandler.RegisterRoutes(v1, lendingService)
}
