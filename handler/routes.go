package handler

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (h *httpHandler) initRoutes(e *echo.Echo) {
	e.GET("/health", h.HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/v1")

	authV1 := v1.Group("/auth")
	authV1.POST("/signup", h.SignUp)

	userV1 := v1.Group("/user")

	patientV1 := v1.Group("/patient")

	hospitalV1 := v1.Group("/hospital")

	wardV1 := v1.Group("/ward")

	roomV1 := v1.Group("/room")

	_ = userV1
	_ = patientV1
	_ = hospitalV1
	_ = wardV1
	_ = roomV1
}
