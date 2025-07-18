package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		c.Next()

		logrus.WithFields(logrus.Fields{
			"status":  c.Response().StatusCode(),
			"method":  c.Method(),
			"path":    c.Path(),
			"latency": time.Since(start),
		}).Info("Request completed")

		return nil
	}
}
