package handler

import (
	"digital-library-dashboard/internal/user/model"
	"digital-library-dashboard/internal/user/service"
	"digital-library-dashboard/pkg/utils"
	"net/http"

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

// Register godoc
// @Summary Register a new user
// @Description Register a new user with name, email, password, and role
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.RegisterRequest true "User registration data"
// @Success 200 {object} model.User
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/register [post]
func (h *Handler) Register(c *fiber.Ctx) error {
	var body model.RegisterRequest
	if err := c.BodyParser(&body); err != nil {
		return utils.ErrorResponseFunc(c, http.StatusBadRequest, "Invalid request")
	}
	user := model.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Role:     body.Role,
	}
	if err := h.svc.Register(&user); err != nil {
		return utils.ErrorResponseFunc(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponseFunc(c, http.StatusOK, "Success register", user)
}

// Login godoc
// @Summary Login user
// @Description Login user with email and password
// @Tags User
// @Accept json
// @Produce json
// @Param credentials body model.Credentials true "User login credentials"
// @Success 200 {object} utils.SuccessResponse{data=utils.TokenResponse}
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /v1/login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	var cred model.Credentials
	if err := c.BodyParser(&cred); err != nil {
		return utils.ErrorResponseFunc(c, http.StatusBadRequest, "Invalid request")
	}
	token, _, err := h.svc.Login(cred.Email, cred.Password)
	if err != nil {
		return utils.ErrorResponseFunc(c, http.StatusUnauthorized, "Invalid credentials")
	}
	return utils.SuccessResponseFunc(c, http.StatusOK, "Success login", fiber.Map{"token": token})
}
