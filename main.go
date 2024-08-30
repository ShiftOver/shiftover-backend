// Package main is the main entry point for shiftover-backend service
package main

import (
	"os"

	"github.com/labstack/gommon/log"

	"github.com/ShiftOver/shiftover-backend/config"
	"github.com/ShiftOver/shiftover-backend/di"
	_ "github.com/ShiftOver/shiftover-backend/docs"
	"github.com/ShiftOver/shiftover-backend/repository"
)

// @title API Endpoints for shiftover-backend
// @version 1.0
// @description This is the API documentation for shiftover-backend

// @BasePath /server
func main() {
	env := os.Getenv("APP_ENV_STAGE")
	if env == "" {
		env = "dev"
	}

	// Initiaize Config
	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Panicf("error - [main.loadConfig] error loading config: %v", err)
	}

	di.New(di.Config{
		AppConfig:   cfg.AppConfig,
		MongoConfig: cfg.MongoConfig,
		UserRepositoryConfig: repository.UserRepositoryConfig{
			CollectionName: os.Getenv("MONGO_COLLECTION_USER"),
		},
	})
}
