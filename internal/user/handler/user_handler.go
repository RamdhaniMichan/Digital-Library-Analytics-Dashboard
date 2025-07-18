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

func (h *Handler) Register(c *fiber.Ctx) error {
	var body model.RegisterRequest
	if err := c.BodyParser(&body); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}
	user := model.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Role:     body.Role,
	}
	if err := h.svc.Register(&user); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, http.StatusOK, "Success register", user)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var cred model.Credentials
	if err := c.BodyParser(&cred); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}
	token, _, err := h.svc.Login(cred.Email, cred.Password)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
	}
	return utils.SuccessResponse(c, http.StatusOK, "Success login", fiber.Map{"token": token})
}
