package routes

import (
	"github.com/gin-gonic/gin"

	"task-api/handlers"
)

// SetupAuthRoutes - setup auth routes
func SetupAuthRoutes(r *gin.Engine, authHandler *handlers.AuthHandler) {
  auth := r.Group("/auth")
  {
    auth.POST("/login", authHandler.Login)
  }
}
