package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"task-api/db"
	"task-api/utils"
)

// HealthCheck - basic health check
func HealthCheck(c *gin.Context) {
  ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
  defer cancel()

  // Quick ping to database
  dbStatus := "healthy"
  err := db.Client.Ping(ctx, nil)
  if err != nil {
    dbStatus = "unhealthy"
  }

  utils.Success(c, http.StatusOK, "API is running", gin.H{
    "status":   "healthy",
    "version":  "1.0.0",
    "database": dbStatus,
  })
}