package routes

import (
	"vayana-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", controllers.Healthcheck)
}
