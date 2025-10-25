package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"task-api/types"
)

// MockAuthService mocks the AuthService interface
type MockAuthService struct {
  mock.Mock
}

// Login mocks the Login method with correct signature
func (m *MockAuthService) Login(ctx context.Context, input types.LoginInput) (*types.LoginResponse, error) {
  // ↑ Ganti interface{} jadi context.Context
  args := m.Called(ctx, input)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*types.LoginResponse), args.Error(1)
}

// setupRouter creates a test router in test mode
func setupRouter() *gin.Engine {
  gin.SetMode(gin.TestMode)
  return gin.Default()
}

// TestMain sets up test environment
func TestMain(m *testing.M) {
  os.Setenv("GO_ENV", "test")
  os.Setenv("JWT_SECRET", "test-secret-key-for-testing-only")
  
  code := m.Run()
  
  os.Unsetenv("JWT_SECRET")
  os.Unsetenv("GO_ENV")
  
  os.Exit(code)
}

func TestAuthHandler_Login(t *testing.T) {
  t.Run("should return 200 on successful login", func(t *testing.T) {
    mockService := new(MockAuthService)
    handler := NewAuthHandler(mockService)
    router := setupRouter()
    router.POST("/auth/login", handler.Login)

    expectedResponse := &types.LoginResponse{
      Token: "test-jwt-token",
      User: types.UserResponse{
        Email: "test@test.com",
        Name:  "Test User",
      },
    }

    mockService.On("Login", mock.Anything, mock.Anything).Return(expectedResponse, nil)

    body := map[string]string{
      "email":    "test@test.com",
      "password": "password123",
    }
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)

    assert.Equal(t, "success", response["status"])
    assert.NotNil(t, response["data"])

    mockService.AssertExpectations(t)
  })

  t.Run("should return 401 on invalid credentials", func(t *testing.T) {
    mockService := new(MockAuthService)
    handler := NewAuthHandler(mockService)
    router := setupRouter()
    router.POST("/auth/login", handler.Login)

    mockService.On("Login", mock.Anything, mock.Anything).
      Return(nil, errors.New("invalid credentials"))

    body := map[string]string{
      "email":    "test@test.com",
      "password": "wrongpassword",
    }
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)

    mockService.AssertExpectations(t)
  })

  t.Run("should return 400 on empty body", func(t *testing.T) {
    mockService := new(MockAuthService)
    handler := NewAuthHandler(mockService)
    router := setupRouter()
    router.POST("/auth/login", handler.Login)

    req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer([]byte("")))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
  })

  t.Run("should return 400 on missing email", func(t *testing.T) {
    mockService := new(MockAuthService)
    handler := NewAuthHandler(mockService)
    router := setupRouter()
    router.POST("/auth/login", handler.Login)

    body := map[string]string{
      "password": "password123",
    }
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
  })

  t.Run("should return 400 on missing password", func(t *testing.T) {
    mockService := new(MockAuthService)
    handler := NewAuthHandler(mockService)
    router := setupRouter()
    router.POST("/auth/login", handler.Login)

    body := map[string]string{
      "email": "test@test.com",
    }
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
  })

  t.Run("should return 400 on invalid JSON", func(t *testing.T) {
    mockService := new(MockAuthService)
    handler := NewAuthHandler(mockService)
    router := setupRouter()
    router.POST("/auth/login", handler.Login)

    req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer([]byte("{invalid json")))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
  })

  t.Run("should return 400 on invalid email format", func(t *testing.T) {
    mockService := new(MockAuthService)
    handler := NewAuthHandler(mockService)
    router := setupRouter()
    router.POST("/auth/login", handler.Login)

    body := map[string]string{
      "email":    "invalid-email",
      "password": "password123",
    }
    jsonBody, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
  })
}
