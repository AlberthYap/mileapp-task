package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes - combine semua routes
func SetupRoutes(r *gin.Engine) {
  // Setup test routes
  SetupTestRoutes(r)

	// TODO: Auth Routes and handle
}
