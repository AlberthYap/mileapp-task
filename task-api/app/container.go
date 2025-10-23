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
  TaskRepo repositories.TaskRepository

  // Services
  AuthService services.AuthService
  TaskService services.TaskService

  // Handlers
  AuthHandler   *handlers.AuthHandler
  TaskHandler   *handlers.TaskHandler
}

// NewContainer - initialize all dependencies
func NewContainer(db *mongo.Database) *Container {
  // Initialize repositories
  userRepo := repositories.NewUserRepository(db)
  taskRepo := repositories.NewTaskRepository(db)

  // Initialize services
  authService := services.NewAuthService(userRepo)
  taskService := services.NewTaskService(taskRepo)

  // Initialize handlers
  authHandler := handlers.NewAuthHandler(authService)
  taskHandler := handlers.NewTaskHandler(taskService)

  return &Container{
    UserRepo:    userRepo,
    TaskRepo:    taskRepo,
    AuthService: authService,
    TaskService: taskService,
    AuthHandler: authHandler,
    TaskHandler: taskHandler,
  }
}
