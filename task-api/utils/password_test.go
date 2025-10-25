package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
  t.Run("should hash password successfully", func(t *testing.T) {
    password := "testpassword123"

    hashed, err := HashPassword(password)

    assert.NoError(t, err)
    assert.NotEmpty(t, hashed)
    assert.NotEqual(t, password, hashed)
    assert.True(t, len(hashed) > 50)
  })

  t.Run("should generate different hash for same password", func(t *testing.T) {
    password := "samepassword"

    hash1, _ := HashPassword(password)
    hash2, _ := HashPassword(password)

    assert.NotEqual(t, hash1, hash2)
  })

  t.Run("should handle empty password", func(t *testing.T) {
    hashed, err := HashPassword("")

    assert.NoError(t, err)
    assert.NotEmpty(t, hashed)
  })
}

func TestCheckPassword(t *testing.T) {
  t.Run("should return true for correct password", func(t *testing.T) {
    password := "correctpassword"
    hashed, _ := HashPassword(password)

    result := CheckPassword(hashed, password)

    assert.True(t, result)
  })

  t.Run("should return false for incorrect password", func(t *testing.T) {
    password := "correctpassword"
    hashed, _ := HashPassword(password)

    result := CheckPassword(hashed, "wrongpassword")

    assert.False(t, result)
  })

  t.Run("should return false for empty password", func(t *testing.T) {
    password := "testpassword"
    hashed, _ := HashPassword(password)

    result := CheckPassword(hashed, "")

    assert.False(t, result)
  })

  t.Run("should return false for invalid hash", func(t *testing.T) {
    result := CheckPassword("invalid-hash", "password")

    assert.False(t, result)
  })

  t.Run("should handle case sensitivity", func(t *testing.T) {
    password := "Password123"
    hashed, _ := HashPassword(password)

    result := CheckPassword(hashed, "password123")

    assert.False(t, result)
  })
}
