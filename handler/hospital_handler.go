package handler

import (
	"context"
	"net/http"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/ShiftOver/shiftover-backend/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

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

// InsertHospital is a handler function to insert a new hospital
// @Summary Insert a new hospital
// @Description Insert a new hospital into the database
// @Tags Hospital
// @Accept json
// @Param payload body dto.HospitalEntity true "Hospital Payload"
// @Success 200 {string} string "Hospital inserted successfully"
// @Router /v1/hospital [post]
func (h *httpHandler) InsertHospital(c echo.Context) error {
	ctx := context.Background()

	var payload dto.HospitalEntity
	if err := c.Bind(&payload); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, errors.Wrap(err, "error - [InsertHospital]: unable to bind payload").Error())
	}

	// Validate the payload
	if err := validate.Struct(payload); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, errors.Wrap(err, "error - [InsertHospital]: validation failed").Error())
	}

	err := h.d.Service.InsertHospital(ctx, &payload)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, errors.Wrap(err, "error - [InsertHospital]: unable to insert hospital").Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Hospital inserted successfully")
}
