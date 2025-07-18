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

// @Summary Create Member
// @Description Create a new member
// @Tags Member
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Param member body model.Member true "Member data"
// @Success 201 {object} model.Member
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/members [post]
// @Security BearerAuth
func (h *Handler) Create(c *fiber.Ctx) error {
	var m model.Member
	if err := c.BodyParser(&m); err != nil {
		return utils.ErrorResponseFunc(c, http.StatusBadGateway, "Invalid input")
	}
	err := h.svc.Create(&m)
	if err != nil {
		return utils.ErrorResponseFunc(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponseFunc(c, http.StatusCreated, "Success created member", m)
}

// @Summary Get Member by ID
// @Description Get member by ID
// @Tags Member
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Param id path int true "Member ID"
// @Success 200 {object} model.Member
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/members/{id} [get]
// @Security BearerAuth
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	member, err := h.svc.GetByID(id)
	if err != nil {
		return utils.ErrorResponseFunc(c, http.StatusNotFound, "Member not found")
	}
	return utils.SuccessResponseFunc(c, http.StatusOK, "Success get member", member)
}

// @Summary List Members
// @Description List all members
// @Tags Member
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Success 200 {object} []model.Member
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/members [get]
// @Security BearerAuth
func (h *Handler) List(c *fiber.Ctx) error {
	members, err := h.svc.List()
	if err != nil {
		return utils.ErrorResponseFunc(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponseFunc(c, http.StatusOK, "Success get members", members)
}
