package handler

import (
	"context"
	"net/http"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/ShiftOver/shiftover-backend/pkg/request"
	"github.com/ShiftOver/shiftover-backend/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
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
		return response.ErrResponse(c, http.StatusInternalServerError, errors.Wrap(err, "error - [GetUser]: unable to fetch user").Error())
	}

	return response.SuccessResponse(c, http.StatusOK, user)
}

// InsertUser is a handler function to insert a new user
// @Summary Insert a new user
// @Description Insert a new user into the database
// @Tags User
// @Accept json
// @Param payload body dto.UserEntity true "User Payload"
// @Success 200 {string} string "User inserted successfully"
// @Router /v1/user [post]
func (h *httpHandler) InsertUser(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	var payload dto.UserEntity
	if err := wrapper.Bind(&payload); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, errors.Wrap(err, "error - [InsertUser]: unable to bind payload").Error())
	}

	err := h.d.Service.InsertUser(ctx, payload)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, errors.Wrap(err, "error - [InsertUser]: unable to insert user").Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "User inserted successfully")
}
