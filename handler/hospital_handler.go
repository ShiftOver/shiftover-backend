package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/ShiftOver/shiftover-backend/pkg/response"
)

// GetHospital is a handler function to fetch a hospital by ID
// @Summary Fetch a hospital by ID
// @Description Fetch a hospital from the database by their ID
// @Tags Hospital
// @Param id path string true "Hospital ID"
// @Success 200 {object} dto.HospitalEntity
// @Router /v1/hospital/{id} [get]
func (h *httpHandler) GetHospital(c echo.Context) error {
	ctx := context.Background()
	hospitalID := c.Param("id")

	hospital, err := h.d.Service.GetHospital(ctx, hospitalID)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, errors.Wrap(err, "error - [GetHospital]: unable to fetch hospital").Error())
	}

	return response.SuccessResponse(c, http.StatusOK, hospital)
}

// ListHospital is a handler function to fetch all hospitals
// @Summary Fetch all hospitals
// @Description Fetch all hospitals from the database
// @Tags Hospital
// @Success 200 {array} dto.HospitalEntity
// @Router /v1/hospital [get]
func (h *httpHandler) ListHospital(c echo.Context) error {
	ctx := context.Background()

	hospitals, err := h.d.Service.ListHospital(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, errors.Wrap(err, "error - [ListHospital]: unable to list hospitals").Error())
	}

	return response.SuccessResponse(c, http.StatusOK, hospitals)
}
