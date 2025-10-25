package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"task-api/models"
)

// ========== INPUT DTOs ==========

// CreateTaskInput - untuk POST /tasks
type CreateTaskInput struct {
  Title       string     `json:"title" binding:"required,min=3,max=200"`
  Description string     `json:"description" binding:"max=1000"`
  Status      string     `json:"status" binding:"omitempty,oneof=pending in_progress completed"`
  Priority    string     `json:"priority" binding:"omitempty,oneof=low medium high"`
  DueDate     *time.Time `json:"due_date"`
  Tags        []string   `json:"tags"`
}

// UpdateTaskInput - untuk PUT /tasks/:id
type UpdateTaskInput struct {
  Title       *string    `json:"title" binding:"omitempty,min=3,max=200"`
  Description *string    `json:"description" binding:"omitempty,max=1000"`
  Status      *string    `json:"status" binding:"omitempty,oneof=pending in_progress completed"`
  Priority    *string    `json:"priority" binding:"omitempty,oneof=low medium high"`
  DueDate     *time.Time `json:"due_date"`
  Tags        []string   `json:"tags"`
}

// TaskQueryParams - untuk GET /tasks
type TaskQueryParams struct {
  Status   string `form:"status" binding:"omitempty,oneof=pending in_progress completed"`
  Priority string `form:"priority" binding:"omitempty,oneof=low medium high"`
  Search   string `form:"search"`
  Sort     string `form:"sort" binding:"omitempty,oneof=created_at -created_at due_date -due_date priority -priority title -title"`
  Page     int    `form:"page" binding:"omitempty,min=1"`
  Limit    int    `form:"limit" binding:"omitempty,min=1,max=100"`
}

// ========== OUTPUT DTOs ==========

// TaskResponse - untuk response API
type TaskResponse struct {
  ID          string     `json:"id"`
  UserID      string     `json:"user_id"`
  Title       string     `json:"title"`
  Description string     `json:"description"`
  Status      string     `json:"status"`
  Priority    string     `json:"priority"`
  DueDate     *time.Time `json:"due_date,omitempty"`
  Tags        []string   `json:"tags"`
  CreatedAt   time.Time  `json:"created_at"`
  UpdatedAt   time.Time  `json:"updated_at"`
  CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// TaskListResponse - untuk list dengan pagination
type TaskListResponse struct {
  Tasks []TaskResponse `json:"tasks"`
  Meta  PaginationMeta `json:"meta"`
}

// ========== CONVERTERS ==========

// ToTaskResponse - convert models.Task ke types.TaskResponse
func ToTaskResponse(task *models.Task) TaskResponse {
  return TaskResponse{
    ID:          task.ID.Hex(),
    UserID:      task.UserID.Hex(),
    Title:       task.Title,
    Description: task.Description,
    Status:      task.Status,
    Priority:    task.Priority,
    DueDate:     task.DueDate,
    Tags:        task.Tags,
    CreatedAt:   task.CreatedAt,
    UpdatedAt:   task.UpdatedAt,
    CompletedAt: task.CompletedAt,
  }
}

// ToTaskResponseList - convert []models.Task ke []types.TaskResponse
func ToTaskResponseList(tasks []models.Task) []TaskResponse {
  responses := make([]TaskResponse, len(tasks))
  for i, task := range tasks {
    responses[i] = ToTaskResponse(&task)
  }
  return responses
}

// ToTask - convert CreateTaskInput ke models.Task
func (input *CreateTaskInput) ToTask(userID bson.ObjectID) models.Task {
  now := time.Now()
  
  status := TaskStatusPending
  if input.Status != "" {
    status = input.Status
  }
  
  priority := TaskPriorityMedium
  if input.Priority != "" {
    priority = input.Priority
  }
  
  tags := input.Tags
  if tags == nil {
    tags = []string{}
  }
  
  return models.Task{
    ID:          bson.NewObjectID(),
    UserID:      userID,
    Title:       input.Title,
    Description: input.Description,
    Status:      status,
    Priority:    priority,
    DueDate:     input.DueDate,
    Tags:        tags,
    CreatedAt:   now,
    UpdatedAt:   now,
  }
}
