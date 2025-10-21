package routes

import (
	"task-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupTestRoutes(r *gin.Engine) {
  // Health endpoint
  r.GET("/health", handlers.HealthCheck)
}
