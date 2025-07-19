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

// @Summary GetAllBooks
// @Description Get all books
// @Tags Books
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} utils.SuccessResponse{data=[]model.BookWithCategory, paginate=utils.Paginate}
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/books [get]
// @Security BearerAuth
func (h *Handler) GetAll(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	books, paginate, err := h.svc.GetAll(page, limit)

	if err != nil {
		return utils.ErrorResponseFunc(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponseFunc(c, http.StatusOK, "Success get books", fiber.Map{
		"data":     books,
		"paginate": paginate,
	})
}

// @Summary GetBookByID
// @Description Get book by ID
// @Tags Books
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Param id path int true "Book ID"
// @Success 200 {object} utils.SuccessResponse{data=model.BookWithCategory}
// @Failure 404 {object} utils.ErrorResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /v1/books/{id} [get]
// @Security BearerAuth
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponseFunc(c, http.StatusBadRequest, "Invalid book ID")
	}
	book, err := h.svc.GetByID(id)
	if err != nil {
		return utils.ErrorResponseFunc(c, http.StatusNotFound, "Book not found")
	}
	return utils.SuccessResponseFunc(c, http.StatusOK, "Success get book", book)
}

// @Summary CreateBook
// @Description Create a new book
// @Tags Books
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Param book body model.Book true "Book data"
// @Success 201 {object} utils.SuccessResponse{message=string}
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/books [post]
// @Security BearerAuth
func (h *Handler) Create(c *fiber.Ctx) error {
	var b model.Book
	user := c.Locals("user").(*userModel.UserClaim)
	if err := c.BodyParser(&b); err != nil {
		return utils.ErrorResponseFunc(c, http.StatusBadRequest, "invalid input")
	}

	b.CreatedBy = user.ID
	if err := h.svc.Create(b); err != nil {
		return utils.ErrorResponseFunc(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponseFunc(c, http.StatusCreated, "success created book", nil)
}

// @Summary UpdateBook
// @Description Update an existing book
// @Tags Books
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Param id path int true "Book ID"
// @Param book body model.Book true "Book data"
// @Success 200 {object} utils.SuccessResponse{data=model.Book}
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/books/{id} [put]
// @Security BearerAuth
func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponseFunc(c, http.StatusBadRequest, "Invalid book ID")
	}
	var b model.Book
	if err := c.BodyParser(&b); err != nil {
		return utils.ErrorResponseFunc(c, http.StatusBadRequest, "Invalid input")
	}
	b.ID = id
	if err := h.svc.Update(b); err != nil {
		return utils.ErrorResponseFunc(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponseFunc(c, http.StatusOK, "success updated book", b)
}

// @Summary DeleteBook
// @Description Delete a book by ID
// @Tags Books
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Param id path int true "Book ID"
// @Success 204 {object} utils.SuccessResponse{message=string}
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/books/{id} [delete]
// @Security BearerAuth
func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponseFunc(c, http.StatusBadRequest, "Invalid book ID")
	}
	if err := h.svc.Delete(id); err != nil {
		return utils.ErrorResponseFunc(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponseFunc(c, http.StatusNoContent, "success deleted book", nil)
}
