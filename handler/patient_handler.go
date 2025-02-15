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

// GetPatient is a handler function to fetch a patient by ID
// @Summary Fetch a patient by ID
// @Description Fetch a patient from the database by their ID
// @Tags Patient
// @Param id path string true "Patient ID"
// @Success 200 {object} dto.PatientEntity
// @Router /v1/patient/{id} [get]
func (h *httpHandler) GetPatient(c echo.Context) error {
	ctx := context.Background()
	patientID := c.Param("id")

	patient, err := h.d.Service.GetPatient(ctx, patientID)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, errors.Wrap(err, "error - [GetPatient]: unable to fetch patient").Error())
	}

	return response.SuccessResponse(c, http.StatusOK, patient)
}

// InsertPatient is a handler function to insert a new patient
// @Summary Insert a new patient
// @Description Insert a new patient into the database
// @Tags Patient
// @Accept json
// @Param payload body dto.PatientModel true "Patient Payload"
// @Success 200 {string} string "Patient inserted successfully"
// @Router /v1/patient [post]
func (h *httpHandler) InsertPatient(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	var payload dto.CreatePatientRequest
	if err := wrapper.Bind(&payload); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [InsertPatient]: unable to bind payload: %v", err))
	}

	patientEntity := dto.PatientEntity{
		FirstName:            payload.FirstName,
		LastName:             payload.LastName,
		DateOfBirth:          payload.DateOfBirth,
		HN:                   payload.HN,
		Sex:                  payload.Sex,
		Allergies:            payload.Allergies,
		ProfilePictureURL:    payload.ProfilePictureURL,
		Education:            payload.Education,
		Occupation:           payload.Occupation,
		Height:               payload.Height,
		Weight:               payload.Weight,
		ModeOfArrival:        payload.ModeOfArrival,
		AdmittedForm:         payload.AdmittedForm,
		InitialVitalSigns:    payload.InitialVitalSigns,
		Diagnosis:            payload.Diagnosis,
		ChiefComplaint:       payload.ChiefComplaint,
		PastIllness:          payload.PastIllness,
		PastIllnessHistory:   payload.PastIllnessHistory,
		FamilyIllnessHistory: payload.FamilyIllnessHistory,
		Reactions:            payload.Reactions,
		Tobacco:              payload.Tobacco,
		Alcohol:              payload.Alcohol,
		Drugs:                payload.Drugs,
		Exercise:             payload.Exercise,
		Sleep:                payload.Sleep,
		Information:          payload.Information,
	}

	err := h.d.Service.InsertPatient(ctx, patientEntity)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [InsertPatient]: unable to insert patient: %v", err))
	}

	return response.SuccessResponse(c, http.StatusOK, "Patient inserted successfully")
}
