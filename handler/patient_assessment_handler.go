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

// UpsertPatientAssessment is a handler function to upsert a patient assessment
// @Summary Upsert a patientAssessment
// @Description Upsert a patientAssessment into the database
// @Tags PatientAssessment
// @Accept json
// @Produce json
// @Param patientAssessment body dto.PatientAssessmentEntity true "PatientAssessment"
// @Success 200 {string} string "PatientAssessment upserted successfully"
// @Router /v1/patient/assessment [post]
func (h *httpHandler) UpsertPatientAssessment(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	var payload dto.PatientAssessmentEntity
	if err := wrapper.Bind(&payload); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [UpsertPatientAssessment]: unable to bind payload: %v", err))
	}

	err := h.d.Service.UpsertPatientAssessment(ctx, payload)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, errors.Wrap(err, "error - [UpsertPatientAssessment]: unable to upsert patientAssessment").Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "PatientAssessment upserted successfully")
}
