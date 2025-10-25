package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"

	"task-api/types"
)

// MockTaskService mocks the TaskService interface
type MockTaskService struct {
  mock.Mock
}

func (m *MockTaskService) CreateTask(ctx context.Context, userID bson.ObjectID, input types.CreateTaskInput) (*types.TaskResponse, error) {
  args := m.Called(ctx, userID, input)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*types.TaskResponse), args.Error(1)
}

func (m *MockTaskService) GetTasks(ctx context.Context, userID bson.ObjectID, query types.TaskQueryParams) (*types.TaskListResponse, error) {
  args := m.Called(ctx, userID, query)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*types.TaskListResponse), args.Error(1)
}

func (m *MockTaskService) GetTask(ctx context.Context, taskID bson.ObjectID, userID bson.ObjectID) (*types.TaskResponse, error) {
  args := m.Called(ctx, taskID, userID)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*types.TaskResponse), args.Error(1)
}

func (m *MockTaskService) UpdateTask(ctx context.Context, taskID bson.ObjectID, userID bson.ObjectID, input types.UpdateTaskInput) (*types.TaskResponse, error) {
  args := m.Called(ctx, taskID, userID, input)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*types.TaskResponse), args.Error(1)
}

func (m *MockTaskService) DeleteTask(ctx context.Context, taskID bson.ObjectID, userID bson.ObjectID) error {
  args := m.Called(ctx, taskID, userID)
  return args.Error(0)
}

func TestTaskHandler_CreateTask(t *testing.T) {
  t.Run("should create task successfully", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    userID := bson.NewObjectID()
    router.Use(func(c *gin.Context) {
      c.Set("userID", userID)
      c.Next()
    })
    router.POST("/tasks", handler.CreateTask)

    expectedResponse := &types.TaskResponse{
      ID:          bson.NewObjectID().Hex(),
      Title:       "Test Task",
      Description: "Test Description",
      Status:      "pending",
      Priority:    "high",
    }

    mockService.On("CreateTask", mock.Anything, userID, mock.Anything).Return(expectedResponse, nil)

    body := map[string]interface{}{
      "title":       "Test Task",
      "description": "Test Description",
      "status":      "pending",
      "priority":    "high",
    }
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)

    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, "success", response["status"])
    assert.NotNil(t, response["data"])

    mockService.AssertExpectations(t)
  })

  t.Run("should fail with empty body", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    router.Use(func(c *gin.Context) {
      c.Set("userID", bson.NewObjectID())
      c.Next()
    })
    router.POST("/tasks", handler.CreateTask)

    req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte("")))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
  })

  t.Run("should fail without authentication", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    router.POST("/tasks", handler.CreateTask)

    body := map[string]string{"title": "Task"}
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)
  })

  t.Run("should fail with invalid input", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    router.Use(func(c *gin.Context) {
      c.Set("userID", bson.NewObjectID())
      c.Next()
    })
    router.POST("/tasks", handler.CreateTask)

    body := map[string]string{
      "title": "",
    }
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
  })

  t.Run("should handle service error", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    userID := bson.NewObjectID()
    router.Use(func(c *gin.Context) {
      c.Set("userID", userID)
      c.Next()
    })
    router.POST("/tasks", handler.CreateTask)

    mockService.On("CreateTask", mock.Anything, userID, mock.Anything).
      Return(nil, errors.New("database error"))

    body := map[string]string{
      "title":  "Task",
      "status": "pending",
    }
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusInternalServerError, w.Code)

    mockService.AssertExpectations(t)
  })
}

func TestTaskHandler_GetTasks(t *testing.T) {
  t.Run("should get all tasks", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    userID := bson.NewObjectID()
    router.Use(func(c *gin.Context) {
      c.Set("userID", userID)
      c.Next()
    })
    router.GET("/tasks", handler.GetTasks)

    // ✅ Fixed: Updated struktur sesuai TaskListResponse baru
    expectedResponse := &types.TaskListResponse{
      Tasks: []types.TaskResponse{
        {ID: "1", Title: "Task 1"},
        {ID: "2", Title: "Task 2"},
      },
      Meta: types.PaginationMeta{
        Page:        1,
        Limit:       10,
        Total:       2,
        TotalPages:  1,
        HasNextPage: false,
        HasPrevPage: false,
      },
    }

    mockService.On("GetTasks", mock.Anything, userID, mock.Anything).Return(expectedResponse, nil)

    req, _ := http.NewRequest("GET", "/tasks", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, "success", response["status"])

    // Verify response structure
    data := response["data"].(map[string]interface{})
    assert.NotNil(t, data["tasks"])
    assert.NotNil(t, data["meta"])

    mockService.AssertExpectations(t)
  })

  t.Run("should handle query parameters", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    userID := bson.NewObjectID()
    router.Use(func(c *gin.Context) {
      c.Set("userID", userID)
      c.Next()
    })
    router.GET("/tasks", handler.GetTasks)

    expectedResponse := &types.TaskListResponse{
      Tasks: []types.TaskResponse{},
      Meta: types.PaginationMeta{
        Page:        1,
        Limit:       10,
        Total:       0,
        TotalPages:  0,
        HasNextPage: false,
        HasPrevPage: false,
      },
    }

    mockService.On("GetTasks", mock.Anything, userID, mock.Anything).Return(expectedResponse, nil)

    req, _ := http.NewRequest("GET", "/tasks?status=pending&priority=high", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    mockService.AssertExpectations(t)
  })

  t.Run("should handle pagination", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    userID := bson.NewObjectID()
    router.Use(func(c *gin.Context) {
      c.Set("userID", userID)
      c.Next()
    })
    router.GET("/tasks", handler.GetTasks)

    expectedResponse := &types.TaskListResponse{
      Tasks: []types.TaskResponse{
        {ID: "1", Title: "Task 1"},
      },
      Meta: types.PaginationMeta{
        Page:        2,
        Limit:       10,
        Total:       25,
        TotalPages:  3,
        HasNextPage: true,
        HasPrevPage: true,
      },
    }

    mockService.On("GetTasks", mock.Anything, userID, mock.Anything).Return(expectedResponse, nil)

    req, _ := http.NewRequest("GET", "/tasks?page=2&limit=10", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)

    data := response["data"].(map[string]interface{})
    meta := data["meta"].(map[string]interface{})
    
    assert.Equal(t, float64(2), meta["page"])
    assert.Equal(t, float64(10), meta["limit"])
    assert.Equal(t, float64(25), meta["total"])
    assert.Equal(t, float64(3), meta["total_pages"])
    assert.Equal(t, true, meta["has_next_page"])
    assert.Equal(t, true, meta["has_prev_page"])

    mockService.AssertExpectations(t)
  })
}

func TestTaskHandler_GetTask(t *testing.T) {
  t.Run("should get task by ID", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()
    
    router.Use(func(c *gin.Context) {
      c.Set("userID", userID)
      c.Next()
    })
    router.GET("/tasks/:id", handler.GetTask)

    expectedResponse := &types.TaskResponse{
      ID:    taskID.Hex(),
      Title: "Test Task",
    }

    mockService.On("GetTask", mock.Anything, taskID, userID).Return(expectedResponse, nil)

    req, _ := http.NewRequest("GET", "/tasks/"+taskID.Hex(), nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    mockService.AssertExpectations(t)
  })

  t.Run("should fail with invalid task ID", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    router.Use(func(c *gin.Context) {
      c.Set("userID", bson.NewObjectID())
      c.Next()
    })
    router.GET("/tasks/:id", handler.GetTask)

    req, _ := http.NewRequest("GET", "/tasks/invalid-id", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
  })

  t.Run("should return 404 when task not found", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()
    
    router.Use(func(c *gin.Context) {
      c.Set("userID", userID)
      c.Next()
    })
    router.GET("/tasks/:id", handler.GetTask)

    mockService.On("GetTask", mock.Anything, taskID, userID).
      Return(nil, errors.New("task not found"))

    req, _ := http.NewRequest("GET", "/tasks/"+taskID.Hex(), nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)

    mockService.AssertExpectations(t)
  })
}

func TestTaskHandler_UpdateTask(t *testing.T) {
  t.Run("should update task successfully", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()
    
    router.Use(func(c *gin.Context) {
      c.Set("userID", userID)
      c.Next()
    })
    router.PUT("/tasks/:id", handler.UpdateTask)

    expectedResponse := &types.TaskResponse{
      ID:     taskID.Hex(),
      Title:  "Updated Task",
      Status: "completed",
    }

    mockService.On("UpdateTask", mock.Anything, taskID, userID, mock.Anything).
      Return(expectedResponse, nil)

    body := map[string]string{
      "title":  "Updated Task",
      "status": "completed",
    }
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("PUT", "/tasks/"+taskID.Hex(), bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    mockService.AssertExpectations(t)
  })

  t.Run("should fail with invalid task ID", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    router.Use(func(c *gin.Context) {
      c.Set("userID", bson.NewObjectID())
      c.Next()
    })
    router.PUT("/tasks/:id", handler.UpdateTask)

    body := map[string]string{"title": "Updated"}
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("PUT", "/tasks/invalid-id", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
  })

  t.Run("should return 404 when task not found", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()
    
    router.Use(func(c *gin.Context) {
      c.Set("userID", userID)
      c.Next()
    })
    router.PUT("/tasks/:id", handler.UpdateTask)

    mockService.On("UpdateTask", mock.Anything, taskID, userID, mock.Anything).
      Return(nil, errors.New("task not found"))

    body := map[string]string{"title": "Updated"}
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("PUT", "/tasks/"+taskID.Hex(), bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)

    mockService.AssertExpectations(t)
  })
}

func TestTaskHandler_DeleteTask(t *testing.T) {
  t.Run("should delete task successfully", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()
    
    router.Use(func(c *gin.Context) {
      c.Set("userID", userID)
      c.Next()
    })
    router.DELETE("/tasks/:id", handler.DeleteTask)

    mockService.On("DeleteTask", mock.Anything, taskID, userID).Return(nil)

    req, _ := http.NewRequest("DELETE", "/tasks/"+taskID.Hex(), nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    mockService.AssertExpectations(t)
  })

  t.Run("should fail with invalid task ID", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    router.Use(func(c *gin.Context) {
      c.Set("userID", bson.NewObjectID())
      c.Next()
    })
    router.DELETE("/tasks/:id", handler.DeleteTask)

    req, _ := http.NewRequest("DELETE", "/tasks/invalid-id", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
  })

  t.Run("should return 404 when task not found", func(t *testing.T) {
    mockService := new(MockTaskService)
    handler := NewTaskHandler(mockService)
    router := setupRouter()
    
    userID := bson.NewObjectID()
    taskID := bson.NewObjectID()
    
    router.Use(func(c *gin.Context) {
      c.Set("userID", userID)
      c.Next()
    })
    router.DELETE("/tasks/:id", handler.DeleteTask)

    mockService.On("DeleteTask", mock.Anything, taskID, userID).
      Return(errors.New("task not found"))

    req, _ := http.NewRequest("DELETE", "/tasks/"+taskID.Hex(), nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)

    mockService.AssertExpectations(t)
  })
}
