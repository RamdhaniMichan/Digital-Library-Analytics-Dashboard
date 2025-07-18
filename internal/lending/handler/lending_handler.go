package handler

import (
	"digital-library-dashboard/internal/lending/model"
	"digital-library-dashboard/internal/lending/service"
	"digital-library-dashboard/internal/middleware"
	userModel "digital-library-dashboard/internal/user/model"
	"digital-library-dashboard/pkg/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	svc service.LendingService
}

func RegisterRoutes(r fiber.Router, svc service.LendingService) {
	h := Handler{svc: svc}
	r.Use(middleware.RoleMiddleware("admin"))
	r.Get("/lendings", h.GetAll)
	r.Get("/lendings/:id", h.GetByID)
	r.Post("/lendings", h.Create)
	r.Put("/lendings/:id", h.Update)
	r.Delete("/lendings/:id", h.Delete)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req model.Lending
	user := c.Locals("user").(*userModel.UserClaim)
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}
	req.BorrowedDate = time.Now()
	req.CreatedBy = user.ID
	req.Status = "borrowed"

	if err := h.svc.Create(req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Success created lending", nil)
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	data, err := h.svc.GetAll()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := h.svc.GetByID(id)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Not found")
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "success get lending", data)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req model.Lending
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}
	req.ID = id

	if err := h.svc.Update(req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, http.StatusOK, "Success updated lending", nil)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	return h.svc.Delete(id)
}
