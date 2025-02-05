package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	ConnectionString string
}

func LoadDatabaseConfig() *DBConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &DBConfig{
		ConnectionString: os.Getenv("DATABASE_URL"),
	}
}

func NewPostgresDB() (*gorm.DB, error) {
	config := LoadDatabaseConfig()

	connectionString := config.ConnectionString

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}
