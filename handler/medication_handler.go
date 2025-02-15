package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/ShiftOver/shiftover-backend/pkg/request"
	"github.com/ShiftOver/shiftover-backend/pkg/response"
	"github.com/labstack/echo/v4"
)

// InsertMedication is a handler function to insert a new medication
// @Summary Insert a new medication
// @Description Insert a new medication into the database
// @Tags Medication
// @Accept json
// @Param payload body dto.MedicationEntity true "Medication Payload"
// @Success 200 {string} string "Medication inserted successfully"
// @Router /v1/medication [post]
func (h *httpHandler) InsertMedication(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	var payload dto.MedicationEntity
	if err := wrapper.Bind(&payload); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [InsertMedication]: unable to bind payload: %v", err))
	}

	err := h.d.Service.InsertMedication(ctx, payload)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [InsertMedication]: unable to insert medication: %v", err))
	}

	return response.SuccessResponse(c, http.StatusOK, "Medication inserted successfully")
}

// UpsertMedication is a handler function to upsert a medication
// @Summary Upsert a medication
// @Description Upsert a medication into the database
// @Tags Medication
// @Accept json
// @Param filter body map[string]interface{} true "Filter"
// @Param payload body dto.MedicationEntity true "Medication Payload"
// @Success 200 {string} string "Medication upserted successfully"
// @Router /v1/medication/upsert [post]
func (h *httpHandler) UpsertMedication(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	var filter map[string]interface{}
	if err := wrapper.Bind(&filter); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [UpsertMedication]: unable to bind filter: %v", err))
	}

	var payload dto.MedicationEntity
	if err := wrapper.Bind(&payload); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [UpsertMedication]: unable to bind payload: %v", err))
	}

	err := h.d.Service.UpsertMedication(ctx, filter, payload)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [UpsertMedication]: unable to upsert medication: %v", err))
	}

	return response.SuccessResponse(c, http.StatusOK, "Medication upserted successfully")
}
