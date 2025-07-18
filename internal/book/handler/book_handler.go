package handler

import (
	"digital-library-dashboard/internal/book/model"
	"digital-library-dashboard/internal/book/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, svc service.Service) {
	h := handler{svc: svc}
	r.Get("/books", h.GetAll)
	r.Get("/books/:id", h.GetByID)
	r.Post("/books", h.Create)
	r.Put("/books/:id", h.Update)
	r.Delete("/books/:id", h.Delete)
}

type handler struct {
	svc service.Service
}

func (h *handler) GetAll(c *fiber.Ctx) error {
	books, err := h.svc.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(books)
}

func (h *handler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}
	book, err := h.svc.GetByID(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

func (h *handler) Create(c *fiber.Ctx) error {
	var b model.Book
	if err := c.BodyParser(&b); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := h.svc.Create(b); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(http.StatusCreated)
}

func (h *handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}
	var b model.Book
	if err := c.BodyParser(&b); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	b.ID = id
	if err := h.svc.Update(b); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(http.StatusOK)
}

func (h *handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}
	if err := h.svc.Delete(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(http.StatusNoContent)
}
