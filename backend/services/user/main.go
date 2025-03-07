package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vayana/pkg/auth"
	"vayana/pkg/db"
	"vayana/pkg/logger"
	"vayana/services/user/api/handlers"
	"vayana/services/user/api/routes"
	"vayana/services/user/config"
	"vayana/services/user/internal/models"
	"vayana/services/user/internal/repository"

	"go.uber.org/zap"
)

// BuildInfo contains build-time information
type BuildInfo struct {
	Version   string
	CommitSHA string
	BuildTime string
}

func main() {
	// Build information
	buildInfo := BuildInfo{
		Version:   "1.0.0",
		CommitSHA: "development",
		BuildTime: time.Now().String(),
	}

	// Initialize logger
	log := logger.NewLogger()
	defer log.Sync()

	// Load configuration
	cfg, err := config.LoadUserServiceConfig()
	if err != nil {
		log.Fatal("Failed to load configuration", zap.Error(err))
	}

	if err := validateOAuthConfig(cfg); err != nil {
		log.Fatal("Invalid OAuth configuration", zap.Error(err))
	}

	// Initialize database connection
	db, err := db.NewPostgresDB()
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}

	// Auto migrate database schemas
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database", zap.Error(err))
	}

	// Initialize JWT manager
	jwtManager := auth.NewJWTManager(cfg.JWTSecret)

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userRepo, jwtManager, log, cfg)

	// Setup router
	router := routes.SetupRouter(userHandler, jwtManager)

	// Register additional components
	// routes.RegisterMetrics(router)
	// routes.RegisterSwagger(router)
	// routes.RegisterCustomMiddleware(router)

	// Create HTTP server
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort),
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Info("Starting server",
			zap.String("host", cfg.ServerHost),
			zap.String("port", cfg.ServerPort),
			zap.String("version", buildInfo.Version),
			zap.String("commit", buildInfo.CommitSHA),
			zap.String("build_time", buildInfo.BuildTime),
		)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	// Create context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown", zap.Error(err))
	}

	log.Info("Server exiting")
}

// validateEnvironment checks if all required environment variables are set
func validateEnvironment() error {
	required := []string{
		"JWT_SECRET",
		"DB_PASSWORD",
		"GOOGLE_CLIENT_ID",
		"GOOGLE_CLIENT_SECRET",
		"GOOGLE_REDIRECT_URL",
	}

	for _, env := range required {
		if os.Getenv(env) == "" {
			return fmt.Errorf("required environment variable %s is not set", env)
		}
	}

	return nil
}

func validateOAuthConfig(cfg *config.UserServiceConfig) error {
	if cfg.GoogleClientID == "" {
		return fmt.Errorf("error: Google Client ID is required")
	}

	if cfg.GoogleClientSecret == "" {
		return fmt.Errorf("error: Google Client Secret is required")
	}

	if cfg.GoogleRedirectURL == "" {
		return fmt.Errorf("error: Google Redirect URL is required")
	}

	return nil
}
