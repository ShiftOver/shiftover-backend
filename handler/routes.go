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
	userV1.GET("/:id", h.GetUser)

	patientV1 := v1.Group("/patient")
	patientV1.GET("/:id", h.GetPatient)
	patientV1.POST("", h.InsertPatient)
	patientV1.GET("/chart/:id", h.GetChartReview)
	patientV1.POST("/chart", h.UpsertChartReview)
	patientV1.GET("/assessment/:id", h.GetPatientAssessment)
	patientV1.POST("/assessment", h.UpsertPatientAssessment)

	hospitalV1 := v1.Group("/hospital")
	hospitalV1.GET("", h.ListHospital)
	hospitalV1.GET("/:id", h.GetHospital)
	hospitalV1.GET("", h.ListHospital)
	hospitalV1.POST("", h.InsertHospital)

	medicationV1 := v1.Group("/medication")
	medicationV1.GET("/:id", h.GetMedication)
	medicationV1.POST("", h.InsertMedication)
	medicationV1.POST("", h.UpsertMedication)

	wardV1 := v1.Group("/ward")

	roomV1 := v1.Group("/room")
	roomV1.GET("", h.ListRooms)
	roomV1.GET("/:id", h.GetRoom)
	roomV1.POST("", h.InsertRoom)

	_ = wardV1
}
