package types

import (
	"task-api/models"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// ========== INPUT DTOs ==========

// LoginInput - request body for login
type LoginInput struct {
  Email    string `json:"email" binding:"required,email"`
  Password string `json:"password" binding:"required,min=8"`
}

// ========== OUTPUT DTOs ==========

// JWTClaims - JWT claims structure
type JWTClaims struct {
  UserID bson.ObjectID `json:"user_id"`
  Email  string        `json:"email"`
  jwt.RegisteredClaims
}

// UserResponse - response user
type UserResponse struct {
  Email     string    `json:"email"`
  Name      string    `json:"name"`
}

// LoginResponse - response login with token (final response)
type LoginResponse struct {
  Token string       `json:"token"`
  User  UserResponse `json:"user"`
}

// ========== CONVERTERS ==========

// ToUserResponse - convert models.User to types.UserResponse
func ToUserResponse(user *models.User) UserResponse {
  return UserResponse{
    Email:     user.Email,
    Name:      user.Name,
  }
}
