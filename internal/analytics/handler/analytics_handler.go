package delivery

import (
	"digital-library-dashboard/internal/analytics/service"
	"digital-library-dashboard/pkg/utils"
	"fmt"
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

// @Summary Get Analytics Data
// @Description Retrieve analytics data for the dashboard
// @Tags Analytics
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Success 200 {object}  model.AnalyticsResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/analytics [get]
// @Security BearerAuth
func (h *Handler) GetAnalytics(c *fiber.Ctx) error {
	data, err := h.svc.GetAnalytics()
	fmt.Println("Analytics Data:", data)
	if err != nil {
		return utils.ErrorResponseFunc(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponseFunc(c, http.StatusOK, "success get analytics", data)
}
