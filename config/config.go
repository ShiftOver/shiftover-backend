// Package config provides configuration settings for the server
package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

// Config represents the configuration of the server
type Config struct {
	AppConfig      AppConfig
	MongoConfig    MongoConfig
	FirebaseConfig FirebaseConfig
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

// FirebaseConfig represents the configuration of the Firebase
type FirebaseConfig struct {
	URL               string
	APIKey            string
	ProjectID         string
	StorageBucket     string
	MessagingSenderID string
	AppID             string
	MeasurementID     string
	ServiceAccountID  string
	CredentialsJSON   string

	UserProfileObjectPath string
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
			DBName:   requiredEnv("MONGO_DBNAME"),
			Username: requiredEnv("MONGO_USERNAME"),
			Password: requiredEnv("MONGO_PASSWORD"),
		},
		FirebaseConfig: FirebaseConfig{
			URL:                   os.Getenv("FIREBASE_URL"),
			APIKey:                os.Getenv("FIREBASE_API_KEY"),
			ProjectID:             requiredEnv("FIREBASE_PROJECT_ID"),
			StorageBucket:         requiredEnv("FIREBASE_STORAGE_BUCKET"),
			MessagingSenderID:     os.Getenv("FIREBASE_MESSAGING_SENDER_ID"),
			AppID:                 os.Getenv("FIREBASE_APP_ID"),
			MeasurementID:         os.Getenv("FIREBASE_MEASUREMENT_ID"),
			ServiceAccountID:      os.Getenv("FIREBASE_SERVICE_ACCOUNT_ID"),
			CredentialsJSON:       os.Getenv("FIREBASE_CREDENTIALS_JSON"),
			UserProfileObjectPath: os.Getenv("FIREBASE_USER_PROFILE_OBJECT_PATH"),
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
