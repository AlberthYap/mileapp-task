package handlers

import (
	"context"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"

	"task-api/services"
	"task-api/types"
	"task-api/utils"
)

type TaskHandler struct {
  taskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *TaskHandler {
  return &TaskHandler{
    taskService: taskService,
  }
}

// CreateTask - POST /tasks - Create new task
func (h *TaskHandler) CreateTask(c *gin.Context) {
  var input types.CreateTaskInput

  if err := c.ShouldBindJSON(&input); err != nil {
    if err == io.EOF {
      log.Warn().Str("ip", c.ClientIP()).Msg("Create task with empty body")
      utils.Fail(c, 400, "Request body required", gin.H{"error": "Please provide task details"})
      return
    }
    
    log.Warn().Err(err).Msg("Create task validation failed")
    utils.Fail(c, 400, types.MsgValidationFailed, gin.H{"error": err.Error()})
    return
  }
  
  // Get user ID from JWT middleware
  userID, exists := c.Get("userID")
  if !exists {
    utils.Fail(c, 401, "Unauthorized", nil)
    return
  }
  
  ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
  defer cancel()
  
  log.Info().
    Str("user_id", userID.(bson.ObjectID).Hex()).
    Str("title", input.Title).
    Msg("Creating task")
  
  response, err := h.taskService.CreateTask(ctx, userID.(bson.ObjectID), input)
  if err != nil {
    log.Error().Err(err).Msg("Failed to create task")
    utils.Error(c, 500, types.MsgInternalError, 0, nil)
    return
  }
  
  log.Info().
    Str("task_id", response.ID).
    Str("user_id", userID.(bson.ObjectID).Hex()).
    Msg("Task created successfully")
  
  utils.Success(c, 201, types.MsgTaskCreated, gin.H{"task": response})
}

// GetTasks - GET /tasks
func (h *TaskHandler) GetTasks(c *gin.Context) {
  var query types.TaskQueryParams
  
  if err := c.ShouldBindQuery(&query); err != nil {
    utils.Fail(c, 400, types.MsgValidationFailed, gin.H{"error": err.Error()})
    return
  }
  
  userID, _ := c.Get("userID")
  
  ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
  defer cancel()
  
  response, err := h.taskService.GetTasks(ctx, userID.(bson.ObjectID), query)
  if err != nil {
    log.Error().Err(err).Msg("Failed to get tasks")
    utils.Error(c, 500, types.MsgInternalError, 0, nil)
    return
  }
  
  utils.Success(c, 200, types.MsgTasksRetrieved, response)
}

// GetTask - GET /tasks/:id - Get single task
func (h *TaskHandler) GetTask(c *gin.Context) {
  taskID := c.Param("id")
  
  objectID, err := bson.ObjectIDFromHex(taskID)
  if err != nil {
    utils.Fail(c, 400, "Invalid task ID", gin.H{"error": "Invalid ID format"})
    return
  }
  
  userID, _ := c.Get("userID")
  
  ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
  defer cancel()
  
  response, err := h.taskService.GetTask(ctx, objectID, userID.(bson.ObjectID))
  if err != nil {
    log.Warn().Err(err).Str("task_id", taskID).Msg("Task not found")
    utils.Fail(c, 404, types.MsgTaskNotFound, nil)
    return
  }
  
  utils.Success(c, 200, types.MsgTaskRetrieved, gin.H{"task": response})
}

// UpdateTask - PUT /tasks/:id - Update task
func (h *TaskHandler) UpdateTask(c *gin.Context) {
  taskID := c.Param("id")
  
  objectID, err := bson.ObjectIDFromHex(taskID)
  if err != nil {
    utils.Fail(c, 400, "Invalid task ID", gin.H{"error": "Invalid ID format"})
    return
  }
  
  var input types.UpdateTaskInput
  
  if err := c.ShouldBindJSON(&input); err != nil {
    utils.Fail(c, 400, types.MsgValidationFailed, gin.H{"error": err.Error()})
    return
  }
  
  userID, _ := c.Get("userID")
  
  ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
  defer cancel()
  
  log.Info().
    Str("task_id", taskID).
    Str("user_id", userID.(bson.ObjectID).Hex()).
    Msg("Updating task")
  
  response, err := h.taskService.UpdateTask(ctx, objectID, userID.(bson.ObjectID), input)
  if err != nil {
    log.Error().Err(err).Str("task_id", taskID).Msg("Failed to update task")
    utils.Fail(c, 404, types.MsgTaskNotFound, nil)
    return
  }
  
  log.Info().Str("task_id", taskID).Msg("Task updated successfully")
  
  utils.Success(c, 200, types.MsgTaskUpdated, gin.H{"task": response})
}

// DeleteTask - DELETE /tasks/:id
func (h *TaskHandler) DeleteTask(c *gin.Context) {
  taskID := c.Param("id")
  
  objectID, err := bson.ObjectIDFromHex(taskID)
  if err != nil {
    utils.Fail(c, 400, "Invalid task ID", gin.H{"error": "Invalid ID format"})
    return
  }
  
  userID, _ := c.Get("userID")
  
  ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
  defer cancel()
  
  log.Info().
    Str("task_id", taskID).
    Str("user_id", userID.(bson.ObjectID).Hex()).
    Msg("Deleting task")
  
  err = h.taskService.DeleteTask(ctx, objectID, userID.(bson.ObjectID))
  if err != nil {
    log.Error().Err(err).Str("task_id", taskID).Msg("Failed to delete task")
    utils.Fail(c, 404, types.MsgTaskNotFound, nil)
    return
  }
  
  log.Info().Str("task_id", taskID).Msg("Task deleted successfully")
  
  utils.Success(c, 200, types.MsgTaskDeleted, gin.H{"deleted_id": taskID})
}
