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
// @Success 201 {object} utils.SuccessResponse{data=model.Member}
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
// @Success 200 {object} utils.SuccessResponse{data=model.Member}
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
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} utils.SuccessResponse{data=[]model.Member, paginate=utils.Paginate}
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/members [get]
// @Security BearerAuth
func (h *Handler) List(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	members, paginate, err := h.svc.List(page, limit)
	if err != nil {
		return utils.ErrorResponseFunc(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponseFunc(c, http.StatusOK, "Success get members", fiber.Map{
		"data":     members,
		"paginate": paginate,
	})
}
