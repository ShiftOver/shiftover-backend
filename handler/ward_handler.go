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

// GetWard is a handler function to fetch a ward by ID
// @Summary Fetch a ward by ID
// @Description Fetch a ward from the database by its ID
// @Tags Ward
// @Param id path string true "Ward ID"
// @Success 200 {object} dto.WardEntity
// @Router /v1/ward/{id} [get]
func (h *httpHandler) GetWard(c echo.Context) error {
	ctx := context.Background()
	wardID := c.Param("id")

	ward, err := h.d.Service.GetWard(ctx, wardID)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [handler.GetWard]: unable to fetch ward: %v", err))
	}

	return response.SuccessResponse(c, http.StatusOK, ward)
}

// ListWards is a handler function to fetch all wards
// @Summary Fetch all wards
// @Description Fetch all wards from the database
// @Tags Ward
// @Success 200 {array} dto.WardEntity
// @Router /v1/ward [get]
func (h *httpHandler) ListWards(c echo.Context) error {
	ctx := context.Background()

	wards, err := h.d.Service.ListWards(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [handler.ListWards]: unable to list wards: %v", err))
	}

	return response.SuccessResponse(c, http.StatusOK, wards)
}

// InsertWard is a handler function to insert a new ward
// @Summary Insert a new ward
// @Description Insert a new ward into the database
// @Tags Ward
// @Accept json
// @Param payload body dto.WardEntity true "Ward Payload"
// @Success 200 {string} string "Ward inserted successfully"
// @Router /v1/ward [post]
func (h *httpHandler) InsertWard(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	var payload dto.WardEntity
	if err := wrapper.Bind(&payload); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [handler.InsertWard]: unable to bind payload: %v", err))
	}

	err := h.d.Service.InsertWard(ctx, payload)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [handler.InsertWard]: unable to insert ward: %v", err))
	}

	return response.SuccessResponse(c, http.StatusCreated, "Ward inserted successfully")
}
