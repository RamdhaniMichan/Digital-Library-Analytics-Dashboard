package middleware

import (
	"digital-library-dashboard/internal/user/model"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if tokenStr == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Missing token"})
		}
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
		}
		return c.Next()
	}
}

func RoleMiddleware(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)
		for _, r := range roles {
			if user.Role == r {
				return c.Next()
			}
		}
		return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
	}
}
