package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ShiftOver/shiftover-backend/pkg/response"
	"github.com/labstack/echo/v4"
)

// GetRoom is a handler function to fetch a room by ID
// @Summary Fetch a room by ID
// @Description Fetch a room from the database by their ID
// @Tags Room
// @Param id path string true "Room ID"
// @Success 200 {object} dto.GetRoomResponse
// @Router /v1/room/{id} [get]
func (h *httpHandler) GetRoom(c echo.Context) error {
	ctx := context.Background()
	roomID := c.Param("id")

	room, err := h.d.Service.GetRoom(ctx, roomID)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [handler.GetRoom] unable to get room: %v", err))
	}

	return response.SuccessResponse(c, http.StatusOK, room)
}

// ListRooms is a handler function to fetch all rooms
// @Summary Fetch all rooms
// @Description Fetch all rooms from the database
// @Tags Room
// @Success 200 {array} dto.GetRoomResponse
// @Router /v1/room [get]
func (h *httpHandler) ListRooms(c echo.Context) error {
	ctx := context.Background()

	rooms, err := h.d.Service.ListRooms(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [handler.ListRooms] unable to list rooms: %v", err))
	}

	return response.SuccessResponse(c, http.StatusOK, rooms)
}
