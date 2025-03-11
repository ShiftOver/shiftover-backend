package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/ShiftOver/shiftover-backend/pkg/request"
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

// UpsertChartReview is a handler function to upsert a chart review into the database
// @Summary Upsert a chartReview
// @Description Upsert a chartReview into the database
// @Tags ChartReview
// @Accept json
// @Param id path string true "Patient ID"
// @Param body body dto.UpdateChartReviewRequest true "ChartReview Request"
// @Success 200
// @Router /v1/patient/chart [post]
func (h *httpHandler) UpsertChartReview(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	var payload dto.UpdateChartReviewRequest
	if err := wrapper.Bind(&payload); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [UpsertChartReview]: unable to bind payload: %v", err))
	}

	err := h.d.Service.UpsertChartReview(ctx, payload)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, errors.Wrap(err, "error - [UpsertChartReview]: unable to upsert chartReview").Error())
	}

	return response.SuccessResponse(c, http.StatusOK, nil)
}
