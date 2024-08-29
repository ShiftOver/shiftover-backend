// Package di provides dependency injection for the server
package di

import (
	"context"
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"

	"github.com/ShiftOver/shiftover-backend/config"
	"github.com/ShiftOver/shiftover-backend/handler"
	"github.com/ShiftOver/shiftover-backend/repository"
	"github.com/ShiftOver/shiftover-backend/service"
)

// Config represents the configuration of the service
type Config struct {
	AppConfig               config.AppConfig
	ExampleRepositoryConfig repository.ExampleRepositoryConfig
}

// New injects the dependencies for the server
func New(c Config) {
	ctx := context.Background()

	e := echo.New()
	setupServer(ctx, e, c)

	exampleRepo := repository.NewExampleRepository(c.ExampleRepositoryConfig)

	service := service.New(service.Dependencies{
		ExampleRepository: exampleRepo,
	})

	handler.New(e, handler.Dependencies{
		Service: service,
	})

	// HTTP Listening
	if err := e.Start(":" + c.AppConfig.Port); err != nil && err != http.ErrServerClosed {
		log.Panicf("error - [main.initServer] unable to start server: %v", err)
	}
}
