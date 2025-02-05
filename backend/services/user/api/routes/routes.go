package routes

import (
	"time"
	"vayana/pkg/auth"
	"vayana/pkg/middleware"
	"vayana/services/user/api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler, jwtManager *auth.JWTManager) *gin.Engine {
	// Create router with default middleware
	router := gin.New()

	// Add recovery middleware
	router.Use(gin.Recovery())

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"time":   time.Now().UTC().Format(time.RFC3339),
		})
	})

	// Service metadata
	router.GET("/metadata", userHandler.GetServiceMetadata)

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		// Public routes
		public := v1.Group("/")
		{
			public.POST("/register", userHandler.RegisterUser)
			public.POST("/login", userHandler.LoginUser)
		}

		// Protected routes
		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware(jwtManager))
		{
			// User profile routes
			protected.GET("/profile", userHandler.GetUserProfile)
			protected.PUT("/profile", userHandler.UpdateUserProfile)
		}

		// Admin routes (if needed in the future)
		// admin := v1.Group("/admin")
		// admin.Use(middleware.AuthMiddleware(jwtManager))
		// admin.Use(middleware.RoleMiddleware("admin"))
		{
			// Add admin routes here when needed
			// Example: admin.GET("/users", userHandler.ListUsers)
		}
	}

	return router
}

/*
// RegisterMetrics registers prometheus metrics for the routes
func RegisterMetrics(router *gin.Engine) {
	// Add prometheus middleware to track request metrics
	router.Use(middleware.PrometheusMiddleware())

	// Expose metrics endpoint for prometheus
	router.GET("/metrics", middleware.PrometheusHandler())
}

// RegisterSwagger registers the swagger documentation routes
func RegisterSwagger(router *gin.Engine) {
	// Swagger documentation
	// TODO: Implement swagger documentation
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// RegisterCustomMiddleware adds any custom middleware to the router
func RegisterCustomMiddleware(router *gin.Engine) {
	// Add request ID middleware
	router.Use(middleware.RequestIDMiddleware())

	// Add logging middleware
	router.Use(middleware.LoggingMiddleware())

	// Add timeout middleware
	router.Use(middleware.TimeoutMiddleware(30 * time.Second))

	// Add rate limiting middleware
	router.Use(middleware.RateLimitMiddleware(100, 1*time.Minute)) // 100 requests per minute
}*/
