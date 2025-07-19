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

// @Summary Create Lending
// @Description Create a new lending record
// @Tags Lending
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Param lending body model.Lending true "Lending data"
// @Success 201 {object} utils.SuccessResponse{data=model.Lending}
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/lendings [post]
// @Security BearerAuth
func (h *Handler) Create(c *fiber.Ctx) error {
	var req model.Lending
	user := c.Locals("user").(*userModel.UserClaim)
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponseFunc(c, fiber.StatusBadRequest, "Invalid input")
	}
	req.BorrowedDate = time.Now()
	req.CreatedBy = user.ID
	req.Status = "borrowed"

	if err := h.svc.Create(req); err != nil {
		return utils.ErrorResponseFunc(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponseFunc(c, http.StatusCreated, "Success created lending", nil)
}

// @Summary Get All Lendings
// @Description Get all lending records
// @Tags Lending
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param member_id query int false "Filter by member ID"
// @Param book_id query int false "Filter by book ID"
// @Param status query string false "Filter by status" Enums(borrowed, returned)
// @Param start_date query string false "Filter by start date (YYYY-MM-DD)"
// @Param end_date query string false "Filter by end date (YYYY-MM-DD)"
// @Success 200 {object} utils.SuccessResponse{data=[]model.Lending, paginate=utils.Paginate}
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/lendings [get]
// @Security BearerAuth
func (h *Handler) GetAll(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	var startDate, endDate time.Time
	if sd := c.Query("start_date"); sd != "" {
		if t, err := time.Parse("2006-01-02", sd); err == nil {
			startDate = t
		}
	}
	if ed := c.Query("end_date"); ed != "" {
		if t, err := time.Parse("2006-01-02", ed); err == nil {
			endDate = t
		}
	}

	filter := model.LendingFilter{
		MemberID:  c.QueryInt("member_id", 0),
		BookID:    c.QueryInt("book_id", 0),
		Status:    c.Query("status"),
		StartDate: startDate,
		EndDate:   endDate,
	}

	data, paginate, err := h.svc.GetAll(page, limit, filter)
	if err != nil {
		return utils.ErrorResponseFunc(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponseFunc(c, fiber.StatusOK, "success get lendings", fiber.Map{
		"data":     data,
		"paginate": paginate,
	})
}

// @Summary Get Lending by ID
// @Description Get a lending record by ID
// @Tags Lending
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Param id path int true "Lending ID"
// @Success 200 {object} utils.SuccessResponse{data=model.Lending}
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/lendings/{id} [get]
// @Security BearerAuth
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := h.svc.GetByID(id)
	if err != nil {
		return utils.ErrorResponseFunc(c, fiber.StatusNotFound, "Not found")
	}
	return utils.SuccessResponseFunc(c, fiber.StatusOK, "success get lending", data)
}

// @Summary Update Lending
// @Description Update a lending record by ID
// @Tags Lending
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Param id path int true "Lending ID"
// @Param lending body model.Lending true "Lending data"
// @Success 200 {object} utils.SuccessResponse{data=model.Lending}
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/lendings/{id} [put]
// @Security BearerAuth
func (h *Handler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req model.Lending
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponseFunc(c, fiber.StatusBadRequest, "Invalid input")
	}
	req.ID = id

	if err := h.svc.Update(req); err != nil {
		return utils.ErrorResponseFunc(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponseFunc(c, http.StatusOK, "Success updated lending", nil)
}

// @Summary Delete Lending
// @Description Delete a lending record by ID
// @Tags Lending
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Param id path int true "Lending ID"
// @Success 200 {object} utils.SuccessResponse{data=string}
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/lendings/{id} [delete]
// @Security BearerAuth
func (h *Handler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	return h.svc.Delete(id)
}
