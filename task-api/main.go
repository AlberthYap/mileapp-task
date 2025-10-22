package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"task-api/app"
	"task-api/configs"
	"task-api/db"
	"task-api/routes"
)

func main() {
  // load env
  configs.LoadEnv()

  // connect to database
  mongoURI := configs.GetEnv("MONGO_URI", "mongodb://localhost:27017")
  db.ConnectDB(mongoURI)

  // Initialize container
  container := app.NewContainer(db.DB)

  // Setup Gin
  r := gin.Default()

  // CORS Middleware
  r.Use(func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    if c.Request.Method == "OPTIONS" {
      c.AbortWithStatus(204)
      return
    }
    c.Next()
  })

	routes.SetupRoutes(r, container)

  // Get port
  port := configs.GetEnv("PORT", "8080")

  // Start server
  log.Printf("Server: http://localhost:%s", port)
  r.Run(":" + port)
}
