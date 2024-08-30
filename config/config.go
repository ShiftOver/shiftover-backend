// Package config provides configuration settings for the server
package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

// Config represents the configuration of the server
type Config struct {
	AppConfig   AppConfig
	MongoConfig MongoConfig
}

// AppConfig represents the configuration of the application
type AppConfig struct {
	Name     string
	Port     string
	EnvStage string
}

// MongoConfig represents the configuration of the MongoDB
type MongoConfig struct {
	HostName string
	DBName   string
	Username string
	Password string
}

// LoadConfig loads the configuration from the .env file
func LoadConfig(env string) (Config, error) {
	if env == "" || env == "dev" {
		if err := godotenv.Load(".env", ".env.secrets"); err != nil {
			return Config{}, err
		}
	}

	return Config{
		AppConfig: AppConfig{
			Name:     requiredEnv("APP_NAME"),
			Port:     requiredEnv("APP_PORT"),
			EnvStage: requiredEnv("APP_ENV_STAGE"),
		},
		MongoConfig: MongoConfig{
			HostName: requiredEnv("MONGO_HOST"),
			Username: requiredEnv("MONGO_USERNAME"),
			Password: requiredEnv("MONGO_PASSWORD"),
		},
	}, nil
}

func requiredEnv(env string) string {
	val := os.Getenv(env)
	if val == "" {
		log.Panic("missing required environment variable: " + env)
	}
	return val
}
