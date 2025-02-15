package handler

import (
	"context"
	"net/http"

	"github.com/ShiftOver/shiftover-backend/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// GetPatientAssessment is a handler function to fetch a patient assessment by patient ID
// @Summary Fetch a patientAssessment by ID
// @Description Fetch a patientAssessment from the database by their ID
// @Tags PatientAssessment
// @Param id path string true "Patient ID"
// @Success 200 {object} dto.PatientAssessmentEntity
// @Router /v1/patient/assessment/{id} [get]
func (h *httpHandler) GetPatientAssessment(c echo.Context) error {
	ctx := context.Background()
	patientID := c.Param("id")

	patientAssessment, err := h.d.Service.GetPatientAssessment(ctx, patientID)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, errors.Wrap(err, "error - [GetPatientAssessment]: unable to fetch patientAssessment").Error())
	}

	return response.SuccessResponse(c, http.StatusOK, patientAssessment)
}
