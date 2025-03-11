package handler

import (
	"context"
	"net/http"

	"github.com/ShiftOver/shiftover-backend/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// GetChartReview is a handler function to fetch a chart review by patient ID
// @Summary Fetch a chartReview by ID
// @Description Fetch a chartReview from the database by their ID
// @Tags ChartReview
// @Param id path string true "Patient ID"
// @Success 200 {object} dto.ChartReviewEntity
// @Router /v1/patient/chart/{id} [get]
func (h *httpHandler) GetChartReview(c echo.Context) error {
	ctx := context.Background()
	patientID := c.Param("id")

	chartReview, err := h.d.Service.GetChartReview(ctx, patientID)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, errors.Wrap(err, "error - [GetPatientAssessment]: unable to fetch chartReview").Error())
	}

	return response.SuccessResponse(c, http.StatusOK, chartReview)
}
