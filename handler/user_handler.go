package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ShiftOver/shiftover-backend/pkg/response"
)

// ListUser lists all the users
// @Summary List all users
// @Description List all users
// @Tags User
// @Success 200 {array} dto.UserEntity
// @Router /v1/user [get]
func (h *httpHandler) ListUser(c echo.Context) error {
	ctx := context.Background()
	users, err := h.d.Service.ListUser(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [ListExample] unable to retrieve data: %v", err))
	}

	return response.SuccessResponse(c, http.StatusOK, users)
}
