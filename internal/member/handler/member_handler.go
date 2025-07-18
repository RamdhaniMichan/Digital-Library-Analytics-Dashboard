package handler

import (
	"digital-library-dashboard/internal/member/model"
	"digital-library-dashboard/internal/member/service"
	"digital-library-dashboard/internal/middleware"
	"digital-library-dashboard/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	svc service.Service
}

func RegisterRoutes(r fiber.Router, svc service.Service) {
	h := Handler{svc: svc}
	r.Use(middleware.RoleMiddleware("admin"))
	r.Get("/members", h.List)
	r.Get("/members/:id", h.GetByID)
	r.Post("/members", h.Create)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var m model.Member
	if err := c.BodyParser(&m); err != nil {
		return utils.ErrorResponse(c, http.StatusBadGateway, "Invalid input")
	}
	err := h.svc.Create(&m)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, http.StatusCreated, "Success created member", m)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	member, err := h.svc.GetByID(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Member not found")
	}
	return utils.SuccessResponse(c, http.StatusOK, "Success get member", member)
}

func (h *Handler) List(c *fiber.Ctx) error {
	members, err := h.svc.List()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, http.StatusOK, "Success get members", members)
}
