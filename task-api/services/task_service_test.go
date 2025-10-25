package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"

	"task-api/models"
	"task-api/types"
)

// MockTaskRepository mocks the TaskRepository interface
type MockTaskRepository struct {
  mock.Mock
}

func (m *MockTaskRepository) Create(ctx context.Context, task *models.Task) error {
  args := m.Called(ctx, task)
  return args.Error(0)
}

func (m *MockTaskRepository) FindByID(ctx context.Context, id bson.ObjectID, userID bson.ObjectID) (*models.Task, error) {
  args := m.Called(ctx, id, userID)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskRepository) FindByUserID(ctx context.Context, userID bson.ObjectID, query types.TaskQueryParams) ([]models.Task, int64, error) {
  args := m.Called(ctx, userID, query)
  if args.Get(0) == nil {
    return nil, args.Get(1).(int64), args.Error(2)
  }
  return args.Get(0).([]models.Task), args.Get(1).(int64), args.Error(2)
}

func (m *MockTaskRepository) Update(ctx context.Context, id bson.ObjectID, userID bson.ObjectID, updates bson.M) error {
  args := m.Called(ctx, id, userID, updates)
  return args.Error(0)
}

func (m *MockTaskRepository) Delete(ctx context.Context, id bson.ObjectID, userID bson.ObjectID) error {
  args := m.Called(ctx, id, userID)
  return args.Error(0)
}

func TestTaskService_CreateTask(t *testing.T) {
  t.Run("should create task successfully", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    input := types.CreateTaskInput{
      Title:       "Test Task",
      Description: "Test Description",
      Status:      "pending",
      Priority:    "high",
      Tags:        []string{"test"},
    }

    mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Task")).Return(nil)

    result, err := service.CreateTask(context.Background(), userID, input)

    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "Test Task", result.Title)
    assert.Equal(t, "Test Description", result.Description)
    assert.Equal(t, "pending", result.Status)
    assert.Equal(t, "high", result.Priority)

    mockRepo.AssertExpectations(t)
  })

  t.Run("should handle repository error", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    input := types.CreateTaskInput{
      Title:  "Task",
      Status: "pending",
    }

    mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Task")).
      Return(errors.New("database error"))

    result, err := service.CreateTask(context.Background(), userID, input)

    assert.Error(t, err)
    assert.Nil(t, result)
    assert.Equal(t, "database error", err.Error())

    mockRepo.AssertExpectations(t)
  })

  t.Run("should handle context timeout", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
    defer cancel()
    time.Sleep(2 * time.Nanosecond)

    userID := bson.NewObjectID()
    input := types.CreateTaskInput{
      Title:  "Task",
      Status: "pending",
    }

    mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Task")).
      Return(context.DeadlineExceeded)

    result, err := service.CreateTask(ctx, userID, input)

    assert.Error(t, err)
    assert.Nil(t, result)
  })
}

func TestTaskService_GetTask(t *testing.T) {
  t.Run("should get task successfully", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()

    mockTask := &models.Task{
      ID:          taskID,
      UserID:      userID,
      Title:       "Test Task",
      Description: "Description",
      Status:      "pending",
      Priority:    "high",
      CreatedAt:   time.Now(),
      UpdatedAt:   time.Now(),
    }

    mockRepo.On("FindByID", mock.Anything, taskID, userID).Return(mockTask, nil)

    result, err := service.GetTask(context.Background(), taskID, userID)

    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, taskID.Hex(), result.ID)
    assert.Equal(t, "Test Task", result.Title)

    mockRepo.AssertExpectations(t)
  })

  t.Run("should return error when task not found", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()

    mockRepo.On("FindByID", mock.Anything, taskID, userID).
      Return(nil, errors.New("task not found"))

    result, err := service.GetTask(context.Background(), taskID, userID)

    assert.Error(t, err)
    assert.Nil(t, result)
    assert.Equal(t, "task not found", err.Error())

    mockRepo.AssertExpectations(t)
  })
}

func TestTaskService_GetTasks(t *testing.T) {
  t.Run("should get all tasks with pagination", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    query := types.TaskQueryParams{
      Page:  1,
      Limit: 10,
    }

    mockTasks := []models.Task{
      {ID: bson.NewObjectID(), UserID: userID, Title: "Task 1", Status: "pending"},
      {ID: bson.NewObjectID(), UserID: userID, Title: "Task 2", Status: "completed"},
    }

    mockRepo.On("FindByUserID", mock.Anything, userID, query).
      Return(mockTasks, int64(2), nil)

    result, err := service.GetTasks(context.Background(), userID, query)

    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Len(t, result.Tasks, 2)
    assert.Equal(t, int64(2), result.Meta.Total)
    assert.Equal(t, 1, result.Meta.Page)
    assert.Equal(t, 10, result.Meta.Limit)
    assert.Equal(t, 1, result.Meta.TotalPages)
    assert.False(t, result.Meta.HasNextPage)
    assert.False(t, result.Meta.HasPrevPage)

    mockRepo.AssertExpectations(t)
  })

  t.Run("should calculate pagination correctly", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    query := types.TaskQueryParams{
      Page:  2,
      Limit: 10,
    }

    mockTasks := []models.Task{
      {ID: bson.NewObjectID(), UserID: userID, Title: "Task 11"},
    }

    mockRepo.On("FindByUserID", mock.Anything, userID, query).
      Return(mockTasks, int64(25), nil)

    result, err := service.GetTasks(context.Background(), userID, query)

    assert.NoError(t, err)
    assert.Equal(t, int64(25), result.Meta.Total)
    assert.Equal(t, 2, result.Meta.Page)
    assert.Equal(t, 3, result.Meta.TotalPages)
    assert.True(t, result.Meta.HasNextPage)
    assert.True(t, result.Meta.HasPrevPage)

    mockRepo.AssertExpectations(t)
  })

  t.Run("should use default pagination values", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    query := types.TaskQueryParams{} // No page/limit

    mockRepo.On("FindByUserID", mock.Anything, userID, query).
      Return([]models.Task{}, int64(0), nil)

    result, err := service.GetTasks(context.Background(), userID, query)

    assert.NoError(t, err)
    assert.Equal(t, 1, result.Meta.Page)
    assert.Equal(t, 10, result.Meta.Limit)

    mockRepo.AssertExpectations(t)
  })

  t.Run("should handle repository error", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    query := types.TaskQueryParams{}

    mockRepo.On("FindByUserID", mock.Anything, userID, query).
      Return(nil, int64(0), errors.New("database error"))

    result, err := service.GetTasks(context.Background(), userID, query)

    assert.Error(t, err)
    assert.Nil(t, result)

    mockRepo.AssertExpectations(t)
  })
}

func TestTaskService_UpdateTask(t *testing.T) {
  t.Run("should update task successfully", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()

    title := "Updated Title"
    status := types.TaskStatusInProgress
    input := types.UpdateTaskInput{
      Title:  &title,
      Status: &status,
    }

    mockRepo.On("Update", mock.Anything, taskID, userID, mock.AnythingOfType("bson.M")).
      Return(nil)

    updatedTask := &models.Task{
      ID:     taskID,
      UserID: userID,
      Title:  "Updated Title",
      Status: "in_progress",
    }
    mockRepo.On("FindByID", mock.Anything, taskID, userID).Return(updatedTask, nil)

    result, err := service.UpdateTask(context.Background(), taskID, userID, input)

    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "Updated Title", result.Title)
    assert.Equal(t, "in_progress", result.Status)

    mockRepo.AssertExpectations(t)
  })

  t.Run("should set completed_at when status is completed", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()

    status := types.TaskStatusCompleted
    input := types.UpdateTaskInput{
      Status: &status,
    }

    var capturedUpdates bson.M
    mockRepo.On("Update", mock.Anything, taskID, userID, mock.AnythingOfType("bson.M")).
      Run(func(args mock.Arguments) {
        capturedUpdates = args.Get(3).(bson.M)
      }).
      Return(nil)

    completedTask := &models.Task{
      ID:          taskID,
      UserID:      userID,
      Status:      "completed",
      CompletedAt: timePtr(time.Now()),
    }
    mockRepo.On("FindByID", mock.Anything, taskID, userID).Return(completedTask, nil)

    result, err := service.UpdateTask(context.Background(), taskID, userID, input)

    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "completed", result.Status)
    assert.NotNil(t, capturedUpdates["completed_at"])

    mockRepo.AssertExpectations(t)
  })

  t.Run("should clear completed_at when status changes from completed", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()

    status := types.TaskStatusPending
    input := types.UpdateTaskInput{
      Status: &status,
    }

    var capturedUpdates bson.M
    mockRepo.On("Update", mock.Anything, taskID, userID, mock.AnythingOfType("bson.M")).
      Run(func(args mock.Arguments) {
        capturedUpdates = args.Get(3).(bson.M)
      }).
      Return(nil)

    task := &models.Task{
      ID:     taskID,
      UserID: userID,
      Status: "pending",
    }
    mockRepo.On("FindByID", mock.Anything, taskID, userID).Return(task, nil)

    _, err := service.UpdateTask(context.Background(), taskID, userID, input)

    assert.NoError(t, err)
    assert.Nil(t, capturedUpdates["completed_at"])

    mockRepo.AssertExpectations(t)
  })

  t.Run("should return error when task not found", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()

    title := "Updated"
    input := types.UpdateTaskInput{Title: &title}

    mockRepo.On("Update", mock.Anything, taskID, userID, mock.AnythingOfType("bson.M")).
      Return(errors.New("task not found"))

    result, err := service.UpdateTask(context.Background(), taskID, userID, input)

    assert.Error(t, err)
    assert.Nil(t, result)

    mockRepo.AssertExpectations(t)
  })

  t.Run("should handle partial updates", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()

    priority := types.TaskPriorityHigh
    input := types.UpdateTaskInput{
      Priority: &priority,
    }

    mockRepo.On("Update", mock.Anything, taskID, userID, mock.AnythingOfType("bson.M")).
      Return(nil)

    task := &models.Task{
      ID:       taskID,
      UserID:   userID,
      Priority: "high",
    }
    mockRepo.On("FindByID", mock.Anything, taskID, userID).Return(task, nil)

    result, err := service.UpdateTask(context.Background(), taskID, userID, input)

    assert.NoError(t, err)
    assert.Equal(t, "high", result.Priority)

    mockRepo.AssertExpectations(t)
  })
}

func TestTaskService_DeleteTask(t *testing.T) {
  t.Run("should delete task successfully", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()

    mockRepo.On("Delete", mock.Anything, taskID, userID).Return(nil)

    err := service.DeleteTask(context.Background(), taskID, userID)

    assert.NoError(t, err)

    mockRepo.AssertExpectations(t)
  })

  t.Run("should return error when task not found", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()

    mockRepo.On("Delete", mock.Anything, taskID, userID).
      Return(errors.New("task not found"))

    err := service.DeleteTask(context.Background(), taskID, userID)

    assert.Error(t, err)
    assert.Equal(t, "task not found", err.Error())

    mockRepo.AssertExpectations(t)
  })

  t.Run("should handle repository error", func(t *testing.T) {
    mockRepo := new(MockTaskRepository)
    service := NewTaskService(mockRepo)

    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()

    mockRepo.On("Delete", mock.Anything, taskID, userID).
      Return(errors.New("database error"))

    err := service.DeleteTask(context.Background(), taskID, userID)

    assert.Error(t, err)

    mockRepo.AssertExpectations(t)
  })
}

// Helper function
func timePtr(t time.Time) *time.Time {
  return &t
}
