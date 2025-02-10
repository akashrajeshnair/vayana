package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type UserServiceConfig struct {
	ServerPort     string
	ServerHost     string
	DatabaseURL    string
	JWTSecret      string
	JWTDuration    string
	ServiceName    string
	ServiceVersion string

	// Google OAuth fields
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
}

func LoadUserServiceConfig() (*UserServiceConfig, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &UserServiceConfig{
		ServerPort:         os.Getenv("SERVER_PORT"),
		ServerHost:         os.Getenv("SERVER_HOST"),
		DatabaseURL:        os.Getenv("DATABASE_URL"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		JWTDuration:        os.Getenv("JWT_DURATION"),
		ServiceName:        os.Getenv("SERVICE_NAME"),
		ServiceVersion:     os.Getenv("SERVICE_VERSION"),
		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		GoogleRedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	}

	return config, nil
}
