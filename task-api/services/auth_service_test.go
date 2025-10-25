package services

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"

	"task-api/models"
	"task-api/types"
	"task-api/utils"
)

// MockUserRepository
type MockUserRepository struct {
  mock.Mock
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
  args := m.Called(ctx, email)
  if args.Get(0) == nil {
    return nil, args.Error(1)
  }
  return args.Get(0).(*models.User), args.Error(1)
}

// TestMain sets up environment for all service tests
func TestMain(m *testing.M) {
  // Setup: Configure JWT_SECRET for token generation
  os.Setenv("GO_ENV", "test")
  os.Setenv("JWT_SECRET", "test-secret-key-for-testing-only")
  
  // Run all tests
  code := m.Run()
  
  // Cleanup
  os.Unsetenv("JWT_SECRET")
  os.Unsetenv("GO_ENV")
  
  os.Exit(code)
}

func TestAuthService_Login(t *testing.T) {
  t.Run("should login successfully with correct credentials", func(t *testing.T) {
    mockRepo := new(MockUserRepository)
    service := NewAuthService(mockRepo)

    hashedPassword, _ := utils.HashPassword("password123")
    mockUser := &models.User{
      ID:        bson.NewObjectID(),
      Email:     "test@test.com",
      Name:      "Test User",
      Password:  hashedPassword,
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
    }

    mockRepo.On("FindByEmail", mock.Anything, "test@test.com").Return(mockUser, nil)

    input := types.LoginInput{
      Email:    "test@test.com",
      Password: "password123",
    }

    result, err := service.Login(context.Background(), input)

    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.NotEmpty(t, result.Token)
    assert.Equal(t, "test@test.com", result.User.Email)
    assert.Equal(t, "Test User", result.User.Name)
    
    mockRepo.AssertExpectations(t)
  })

  t.Run("should fail with wrong password", func(t *testing.T) {
    mockRepo := new(MockUserRepository)
    service := NewAuthService(mockRepo)

    hashedPassword, _ := utils.HashPassword("correctpassword")
    mockUser := &models.User{
      ID:       bson.NewObjectID(),
      Email:    "test@test.com",
      Password: hashedPassword,
    }

    mockRepo.On("FindByEmail", mock.Anything, "test@test.com").Return(mockUser, nil)

    input := types.LoginInput{
      Email:    "test@test.com",
      Password: "wrongpassword",
    }

    result, err := service.Login(context.Background(), input)

    assert.Error(t, err)
    assert.Nil(t, result)
    assert.Contains(t, err.Error(), "invalid credentials")

    mockRepo.AssertExpectations(t)
  })

  t.Run("should fail when user not found", func(t *testing.T) {
    mockRepo := new(MockUserRepository)
    service := NewAuthService(mockRepo)

    mockRepo.On("FindByEmail", mock.Anything, "notfound@test.com").
      Return(nil, errors.New("user not found"))

    input := types.LoginInput{
      Email:    "notfound@test.com",
      Password: "password123",
    }

    result, err := service.Login(context.Background(), input)

    assert.Error(t, err)
    assert.Nil(t, result)

    mockRepo.AssertExpectations(t)
  })
}
