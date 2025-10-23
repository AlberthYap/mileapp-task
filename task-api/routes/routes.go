package routes

import (
	"task-api/app"

	"github.com/gin-gonic/gin"
)

// SetupRoutes - setup routes
func SetupRoutes(r *gin.Engine, c *app.Container) {
  // Setup test routes
  SetupTestRoutes(r)

  // Auth routes
  SetupAuthRoutes(r, c.AuthHandler)


  SetupTaskRoutes(r, c.TaskHandler)
}
