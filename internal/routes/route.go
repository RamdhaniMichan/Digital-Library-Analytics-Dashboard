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

	analyticsHandler "digital-library-dashboard/internal/analytics/handler"
	analyticsRepository "digital-library-dashboard/internal/analytics/repository"
	analyticsService "digital-library-dashboard/internal/analytics/service"
)

func SetupRoute(auth fiber.Router, db *sql.DB) {
	ur := userRepository.NewRepository(db)
	us := userService.NewService(ur)
	userHandler.RegisterRoutes(auth.Group("/v1"), us)

	v1 := auth.Group("/v1", middleware.JWTMiddleware())

	br := bookRepository.NewRepository(db)
	bs := bookService.NewService(br)
	bookHandler.RegisterRoutes(v1, bs)

	mr := memberRepository.NewRepository(db)
	ms := memberService.NewService(mr)
	memberHandler.RegisterRoutes(v1, ms)

	lr := lendingRepository.NewLendingRepository(db)
	ls := lendingService.NewLendingService(lr, br, mr)
	lendingHandler.RegisterRoutes(v1, ls)

	ar := analyticsRepository.NewAnalyticsRepository(db)
	as := analyticsService.NewAnalyticsService(ar)
	analyticsHandler.RegisterRoutes(v1, as)
}
