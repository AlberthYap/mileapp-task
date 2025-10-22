package app

import (
	"go.mongodb.org/mongo-driver/v2/mongo"

	"task-api/handlers"
	"task-api/repositories"
	"task-api/services"
)

// Container - holds all dependencies
type Container struct {
  // Repositories
  UserRepo repositories.UserRepository

  // Services
  AuthService services.AuthService

  // Handlers
  AuthHandler   *handlers.AuthHandler
}

// NewContainer - initialize all dependencies
func NewContainer(db *mongo.Database) *Container {
  // Initialize repositories
  userRepo := repositories.NewUserRepository(db)

  // Initialize services
  authService := services.NewAuthService(userRepo)

  // Initialize handlers
  authHandler := handlers.NewAuthHandler(authService)

  return &Container{
    UserRepo:      userRepo,
    AuthService:   authService,
    AuthHandler:   authHandler,
  }
}
