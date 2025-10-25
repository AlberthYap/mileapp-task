package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"

	"task-api/utils"
)

// TestMain sets up environment for middleware tests
func TestMain(m *testing.M) {
  os.Setenv("GO_ENV", "test")
  os.Setenv("JWT_SECRET", "test-secret-key-for-testing-only")
  
  code := m.Run()
  
  os.Unsetenv("JWT_SECRET")
  os.Unsetenv("GO_ENV")
  
  os.Exit(code)
}

func TestAuthMiddleware(t *testing.T) {
  gin.SetMode(gin.TestMode)

  t.Run("should pass with valid token", func(t *testing.T) {
    // Setup
    router := gin.New()
    router.Use(AuthMiddleware())
    router.GET("/protected", func(c *gin.Context) {
      userID := c.GetString("userID")
      userEmail := c.GetString("userEmail")
      c.JSON(200, gin.H{
        "userID":    userID,
        "userEmail": userEmail,
        "message":   "success",
      })
    })

    // Generate valid token
    userID := bson.NewObjectID()
    email := "test@test.com"
    token, _ := utils.GenerateToken(userID, email)

    // Execute
    req := httptest.NewRequest("GET", "/protected", nil)
    req.Header.Set("Authorization", "Bearer "+token)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Assert
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "success")
    assert.Contains(t, w.Body.String(), email)
  })

  t.Run("should fail with missing authorization header", func(t *testing.T) {
    router := gin.New()
    router.Use(AuthMiddleware())
    router.GET("/protected", func(c *gin.Context) {
      c.JSON(200, gin.H{"message": "success"})
    })

    req := httptest.NewRequest("GET", "/protected", nil)
    // No Authorization header
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)
    assert.Contains(t, w.Body.String(), "Authorization header required")
  })

  t.Run("should fail with invalid token format", func(t *testing.T) {
    router := gin.New()
    router.Use(AuthMiddleware())
    router.GET("/protected", func(c *gin.Context) {
      c.JSON(200, gin.H{"message": "success"})
    })

    req := httptest.NewRequest("GET", "/protected", nil)
    req.Header.Set("Authorization", "Bearer invalid-token")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)
    assert.Contains(t, w.Body.String(), "Invalid or expired token")
  })

  t.Run("should fail with malformed token", func(t *testing.T) {
    router := gin.New()
    router.Use(AuthMiddleware())
    router.GET("/protected", func(c *gin.Context) {
      c.JSON(200, gin.H{"message": "success"})
    })

    req := httptest.NewRequest("GET", "/protected", nil)
    req.Header.Set("Authorization", "InvalidFormat")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)
  })

  t.Run("should handle token without Bearer prefix", func(t *testing.T) {
    router := gin.New()
    router.Use(AuthMiddleware())
    router.GET("/protected", func(c *gin.Context) {
      c.JSON(200, gin.H{"message": "success"})
    })

    // Generate valid token
    userID := bson.NewObjectID()
    token, _ := utils.GenerateToken(userID, "test@test.com")

    req := httptest.NewRequest("GET", "/protected", nil)
    req.Header.Set("Authorization", token) // Without "Bearer "
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
  })

  t.Run("should fail with empty token", func(t *testing.T) {
    router := gin.New()
    router.Use(AuthMiddleware())
    router.GET("/protected", func(c *gin.Context) {
      c.JSON(200, gin.H{"message": "success"})
    })

    req := httptest.NewRequest("GET", "/protected", nil)
    req.Header.Set("Authorization", "Bearer ")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)
  })

  t.Run("should set user context correctly", func(t *testing.T) {
    router := gin.New()
    router.Use(AuthMiddleware())
    router.GET("/protected", func(c *gin.Context) {
      // Check if context is set correctly
      userID, exists := c.Get("userID")
      assert.True(t, exists)
      assert.NotNil(t, userID)

      userEmail, exists := c.Get("userEmail")
      assert.True(t, exists)
      assert.Equal(t, "test@test.com", userEmail)

      c.JSON(200, gin.H{"message": "success"})
    })

    userID := bson.NewObjectID()
    token, _ := utils.GenerateToken(userID, "test@test.com")

    req := httptest.NewRequest("GET", "/protected", nil)
    req.Header.Set("Authorization", "Bearer "+token)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
  })
}
