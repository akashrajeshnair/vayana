package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type UserServiceConfig struct {
	ServerPort     string
	ServerHost     string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	JWTSecret      string
	JWTDuration    string
	ServiceName    string
	ServiceVersion string
}

func LoadUserServiceConfig() (*UserServiceConfig, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &UserServiceConfig{
		ServerPort:     os.Getenv("SERVER_PORT"),
		ServerHost:     os.Getenv("SERVER_HOST"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		JWTDuration:    os.Getenv("JWT_DURATION"),
		ServiceName:    os.Getenv("SERVICE_NAME"),
		ServiceVersion: os.Getenv("SERVICE_VERSION"),
	}

	return config, nil
}

func (c *UserServiceConfig) GetPostgresURL() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost,
		c.DBPort,
		c.DBUser,
		c.DBPassword,
		c.DBName,
	)
}
