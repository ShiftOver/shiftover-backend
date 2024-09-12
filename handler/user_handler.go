package handler

import (
	"context"
	"net/http"

	"github.com/ShiftOver/shiftover-backend/pkg/response"
	"github.com/labstack/echo/v4"
)

// GetUser is a handler function to fetch a user by ID
// @Summary Fetch a user by ID
// @Description Fetch a user from the database by their ID
// @Tags User
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserEntity
// @Router /v1/user/{id} [get]
func (h *httpHandler) GetUser(c echo.Context) error {
	ctx := context.Background()
	userID := c.Param("id")

	user, err := h.d.Service.GetUser(ctx, userID)
	if err != nil {
		return response.ErrResponse(c, http.StatusNotFound, "User not found")
	}

	return response.SuccessResponse(c, http.StatusOK, user)
}
