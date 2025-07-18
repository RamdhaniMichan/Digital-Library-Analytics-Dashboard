package handler

import (
	"digital-library-dashboard/internal/book/model"
	"digital-library-dashboard/internal/book/service"
	"digital-library-dashboard/internal/middleware"
	userModel "digital-library-dashboard/internal/user/model"
	"digital-library-dashboard/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, svc service.Service) {
	h := Handler{svc: svc}
	r.Use(middleware.RoleMiddleware("admin"))
	r.Get("/books", h.GetAll)
	r.Get("/books/:id", h.GetByID)
	r.Post("/books", h.Create)
	r.Put("/books/:id", h.Update)
	r.Delete("/books/:id", h.Delete)
}

type Handler struct {
	svc service.Service
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	books, err := h.svc.GetAll()
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, http.StatusOK, "Success get books", books)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid book ID")
	}
	book, err := h.svc.GetByID(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Book not found")
	}
	return utils.SuccessResponse(c, http.StatusOK, "Success get book", book)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var b model.Book
	user := c.Locals("user").(*userModel.UserClaim)
	if err := c.BodyParser(&b); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "invalid input")
	}

	b.CreatedBy = user.ID
	if err := h.svc.Create(b); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, http.StatusCreated, "success created book", nil)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid book ID")
	}
	var b model.Book
	if err := c.BodyParser(&b); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
	}
	b.ID = id
	if err := h.svc.Update(b); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, http.StatusOK, "success updated book", b)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid book ID")
	}
	if err := h.svc.Delete(id); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, http.StatusNoContent, "success deleted book", nil)
}
