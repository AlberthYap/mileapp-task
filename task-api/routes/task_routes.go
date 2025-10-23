package routes

import (
	"github.com/gin-gonic/gin"

	"task-api/handlers"
	"task-api/middleware"
)

func SetupTaskRoutes(r *gin.Engine, taskHandler *handlers.TaskHandler) {
  tasks := r.Group("/tasks")
  tasks.Use(middleware.AuthMiddleware()) // Protected routes
  {
    tasks.POST("", taskHandler.CreateTask)      // Create task
    tasks.GET("", taskHandler.GetTasks)         // Get all tasks (with filters)
    tasks.GET("/:id", taskHandler.GetTask)      // Get single task
    tasks.PUT("/:id", taskHandler.UpdateTask)   // Update task
    tasks.DELETE("/:id", taskHandler.DeleteTask) // Delete task
  }
}
