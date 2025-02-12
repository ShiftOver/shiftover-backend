package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/ShiftOver/shiftover-backend/pkg/request"
	"github.com/ShiftOver/shiftover-backend/pkg/response"
)

// SignUp is a handler function to create a new user
// @Summary Create a new user
// @Description Create a new user in firebase and users mongo collection
// @Tags Auth
// @Success 200 {array} dto.SignUpRes
// @Router /v1/auth/signup [post]
func (h *httpHandler) SignUp(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(dto.SignUpReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [handler.SignUp] unable to bind request: %v", err))
	}

	res, err := h.d.Service.SignUp(ctx, *req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [handler.SignUp] unable to sign up user: %v", err))
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
