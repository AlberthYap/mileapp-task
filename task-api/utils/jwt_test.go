package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestMain(m *testing.M) {
  os.Setenv("GO_ENV", "test")
  os.Setenv("JWT_SECRET", "test-secret-key-for-testing-only")
  
  code := m.Run()
  
  os.Unsetenv("JWT_SECRET")
  os.Unsetenv("GO_ENV")
  
  os.Exit(code)
}

func TestGenerateToken(t *testing.T) {
  t.Run("should generate token successfully", func(t *testing.T) {
    userID := bson.NewObjectID()
    email := "test@test.com"

    token, err := GenerateToken(userID, email)

    assert.NoError(t, err)
    assert.NotEmpty(t, token)
    assert.True(t, len(token) > 20)
  })

  t.Run("should generate different tokens", func(t *testing.T) {
    userID1 := bson.NewObjectID()
    userID2 := bson.NewObjectID()

    token1, _ := GenerateToken(userID1, "test@test.com")
    token2, _ := GenerateToken(userID2, "test@test.com")

    assert.NotEqual(t, token1, token2)
  })
}

func TestValidateToken(t *testing.T) {
  t.Run("should validate correct token", func(t *testing.T) {
    userID := bson.NewObjectID()
    email := "test@test.com"
    token, _ := GenerateToken(userID, email)

    claims, err := ValidateToken(token)

    assert.NoError(t, err)
    assert.NotNil(t, claims)
    assert.Equal(t, userID, claims.UserID)
    assert.Equal(t, email, claims.Email)
  })

  t.Run("should reject invalid token", func(t *testing.T) {
    claims, err := ValidateToken("invalid.token")

    assert.Error(t, err)
    assert.Nil(t, claims)
  })
}
