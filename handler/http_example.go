package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ShiftOver/shiftover-backend/pkg/response"
)

// DoExample handles the request to do example function
// @Summary Run DoExample function
// @Description Run DoExample function
// @Success 200 string string
// @Router /v1/example [get]
func (h *httpHandler) DoExample(c echo.Context) error {
	ctx := context.Background()
	examples, err := h.d.Service.DoExample(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [DoExample] unable to retrieve data: %v", err))
	}
	return response.SuccessResponse(c, http.StatusOK, examples)
}
