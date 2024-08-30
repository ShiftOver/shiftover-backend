// Package di provides dependency injection for the server
package di

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/ShiftOver/shiftover-backend/config"
	"github.com/ShiftOver/shiftover-backend/handler"
	"github.com/ShiftOver/shiftover-backend/repository"
	"github.com/ShiftOver/shiftover-backend/service"
)

// Config represents the configuration of the service
type Config struct {
	AppConfig            config.AppConfig
	MongoConfig          config.MongoConfig
	FirebaseConfig       config.FirebaseConfig
	UserRepositoryConfig repository.UserRepositoryConfig
}

// New injects the dependencies for the server
func New(c Config) {
	ctx := context.Background()

	e := echo.New()
	setupServer(ctx, e, c)

	authRepo, storageRepo, err := setupFirebase(ctx, c.FirebaseConfig)
	if err != nil {
		log.Panicf("error - [di.setupFirebase] unable to initialize Firebase client: %v", err)
	}

	userRepo := newMongoRepositories(ctx, c)

	service := service.New(service.Dependencies{
		UserRepository:    userRepo,
		AuthRepository:    authRepo,
		StorageRepository: storageRepo,
	})

	handler.New(e, handler.Dependencies{
		Service: service,
	})

	// HTTP Listening
	if err := e.Start(":" + c.AppConfig.Port); err != nil && err != http.ErrServerClosed {
		log.Panicf("error - [main.initServer] unable to start server: %v", err)
	}
}
