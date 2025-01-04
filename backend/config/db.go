package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Database *sql.DB

func ConnectDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file: " + err.Error())
	}

	connectionString := os.Getenv("DATABASE_URL")

	Database, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	err = Database.Ping()
	if err != nil {
		log.Fatal("Database offline... Ping Failed: ", err)
	}

	log.Println("Database connected!")
	return Database
}

func MigrateDatabase() error {
	files, err := filepath.Glob("migrations/*.sql")
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
	for _, file := range files {
		log.Println("Migrating: ", file)
		operations, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}

		_, err = Database.Exec(string(operations))
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", file, err)
		}
	}
	log.Println("Database migration completed successfully!")
	return nil
}

func CloseDatabase() {
	if Database != nil {
		Database.Close()
	}
}
