package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"task-api/app"
	"task-api/configs"
	"task-api/db"
	"task-api/middleware"
	"task-api/routes"
	"task-api/utils"
)

func main() {
  // Initialize logger
  utils.InitLogger()

  // load env
  configs.LoadEnv()

  // connect to database
  mongoURI := configs.GetEnv("MONGO_URI", "mongodb://localhost:27017")
  db.ConnectDB(mongoURI)

  // Initialize container
  container := app.NewContainer(db.DB)

  // Setup Gin
  r := gin.New()

  // Logger Middleware
  r.Use(middleware.LoggerMiddleware())

  // Setup trusted proxies
  utils.SetupTrustedProxies(r)

  // CORS Middleware
  r.Use(middleware.CORSMiddleware())

	routes.SetupRoutes(r, container)

  // Get port
  port := configs.GetEnv("PORT", "8080")
  log.Info().Str("port", port).Msg("Server starting")

  // Start server
  log.Printf("Server: http://localhost:%s", port)
  r.Run(":" + port)
}
