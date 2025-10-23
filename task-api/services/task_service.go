package services

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"task-api/repositories"
	"task-api/types"
)

// TaskService - interface
type TaskService interface {
  CreateTask(ctx context.Context, userID bson.ObjectID, input types.CreateTaskInput) (*types.TaskResponse, error)
  GetTask(ctx context.Context, taskID bson.ObjectID, userID bson.ObjectID) (*types.TaskResponse, error)
  GetTasks(ctx context.Context, userID bson.ObjectID, query types.TaskQueryParams) (*types.TaskListResponse, error)
  UpdateTask(ctx context.Context, taskID bson.ObjectID, userID bson.ObjectID, input types.UpdateTaskInput) (*types.TaskResponse, error)
  DeleteTask(ctx context.Context, taskID bson.ObjectID, userID bson.ObjectID) error
}

// taskService - implementation
type taskService struct {
  taskRepo repositories.TaskRepository
}

// NewTaskService - constructor
func NewTaskService(taskRepo repositories.TaskRepository) TaskService {
  return &taskService{
    taskRepo: taskRepo,
  }
}

// CreateTask - create new task
func (s *taskService) CreateTask(ctx context.Context, userID bson.ObjectID, input types.CreateTaskInput) (*types.TaskResponse, error) {
  ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
  defer cancel()
  
  // Convert input to model
  task := input.ToTask(userID)
  
  // Save to database
  err := s.taskRepo.Create(ctx, &task)
  if err != nil {
    return nil, err
  }
  
  // Convert to response
  response := types.ToTaskResponse(&task)
  return &response, nil
}

// GetTask - get single task
func (s *taskService) GetTask(ctx context.Context, taskID bson.ObjectID, userID bson.ObjectID) (*types.TaskResponse, error) {
  ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
  defer cancel()
  
  task, err := s.taskRepo.FindByID(ctx, taskID, userID)
  if err != nil {
    return nil, err
  }
  
  response := types.ToTaskResponse(task)
  return &response, nil
}

// GetTasks - get tasks with filter
func (s *taskService) GetTasks(ctx context.Context, userID bson.ObjectID, query types.TaskQueryParams) (*types.TaskListResponse, error) {
  ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
  defer cancel()
  
  tasks, total, err := s.taskRepo.FindByUserID(ctx, userID, query)
  if err != nil {
    return nil, err
  }
  
  // Convert to response
  taskResponses := types.ToTaskResponseList(tasks)
  
  // Pagination meta
  page := 1
  if query.Page > 0 {
    page = query.Page
  }
  
  limit := 10
  if query.Limit > 0 {
    limit = query.Limit
  }
  
  totalPages := int(total) / limit
  if int(total)%limit != 0 {
    totalPages++
  }
  
  meta := types.PaginationMeta{
    Page:        page,
    Limit:       limit,
    Total:       total,
    TotalPages:  totalPages,
    HasNextPage: page < totalPages,
    HasPrevPage: page > 1,
  }
  
  return &types.TaskListResponse{
    Tasks: taskResponses,
    Meta:  meta,
  }, nil
}

// UpdateTask - update task
func (s *taskService) UpdateTask(ctx context.Context, taskID bson.ObjectID, userID bson.ObjectID, input types.UpdateTaskInput) (*types.TaskResponse, error) {
  ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
  defer cancel()
  
  // Build update document
  updates := bson.M{}
  
  if input.Title != nil {
    updates["title"] = *input.Title
  }
  
  if input.Description != nil {
    updates["description"] = *input.Description
  }
  
  if input.Status != nil {
    updates["status"] = *input.Status
    
    // If status is completed, set completed_at
    if *input.Status == types.TaskStatusCompleted {
      now := time.Now()
      updates["completed_at"] = now
    } else {
      updates["completed_at"] = nil
    }
  }
  
  if input.Priority != nil {
    updates["priority"] = *input.Priority
  }
  
  if input.DueDate != nil {
    updates["due_date"] = input.DueDate
  }
  
  if input.Tags != nil {
    updates["tags"] = input.Tags
  }
  
  // Update
  err := s.taskRepo.Update(ctx, taskID, userID, updates)
  if err != nil {
    return nil, err
  }
  
  // Get updated task
  task, err := s.taskRepo.FindByID(ctx, taskID, userID)
  if err != nil {
    return nil, err
  }
  
  response := types.ToTaskResponse(task)
  return &response, nil
}

// DeleteTask - delete task
func (s *taskService) DeleteTask(ctx context.Context, taskID bson.ObjectID, userID bson.ObjectID) error {
  ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
  defer cancel()
  
  return s.taskRepo.Delete(ctx, taskID, userID)
}
