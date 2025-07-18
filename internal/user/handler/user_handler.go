package handler

import (
	"digital-library-dashboard/internal/user/model"
	"digital-library-dashboard/internal/user/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	svc service.Service
}

func RegisterRoutes(r fiber.Router, svc service.Service) {
	h := Handler{svc: svc}
	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var body model.RegisterRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	user := model.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Role:     body.Role,
	}
	if err := h.svc.Register(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var cred model.Credentials
	if err := c.BodyParser(&cred); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	token, role, err := h.svc.Login(cred.Email, cred.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}
	return c.JSON(fiber.Map{"token": token, "role": role})
}
