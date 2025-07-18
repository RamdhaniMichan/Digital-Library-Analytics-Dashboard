package delivery

import (
	"digital-library-dashboard/internal/analytics/service"
	"digital-library-dashboard/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	svc service.AnalyticsService
}

func RegisterRoutes(r fiber.Router, svc service.AnalyticsService) {
	h := Handler{svc: svc}
	r.Get("/analytics", h.GetAnalytics)
}

func (h *Handler) GetAnalytics(c *fiber.Ctx) error {
	data, err := h.svc.GetAnalytics(c.Context())
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, http.StatusOK, "success get analytics", data)
}
