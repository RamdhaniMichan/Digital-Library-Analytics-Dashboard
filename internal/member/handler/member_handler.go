package handler

import (
	"digital-library-dashboard/internal/member/model"
	"digital-library-dashboard/internal/member/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	svc service.Service
}

func RegisterRoutes(r fiber.Router, svc service.Service) {
	h := Handler{svc: svc}

	r.Get("/members", h.List)
	r.Get("/members/:id", h.GetByID)
	r.Post("/members", h.Create)
	// r.Put("/members/:id", h.Update)
	// r.Delete("/members/:id", h.Delete)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var m model.Member
	if err := c.BodyParser(&m); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	err := h.svc.Create(&m)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(m)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	member, err := h.svc.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "member not found"})
	}
	return c.JSON(member)
}

func (h *Handler) List(c *fiber.Ctx) error {
	members, err := h.svc.List()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(members)
}
