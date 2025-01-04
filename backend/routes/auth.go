package routes

import (
	"database/sql"
	"vayana-backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, db *sql.DB) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", controllers.Register(db))
		authGroup.POST("/login", controllers.Login(db))
	}
}
