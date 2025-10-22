package services

import (
	"context"
	"errors"
	"time"

	"task-api/repositories"
	"task-api/types"
	"task-api/utils"
)

// AuthService - interface for auth service
type AuthService interface {
  Login(ctx context.Context, input types.LoginInput) (*types.LoginResponse, error)
}

// authService - implement AuthService
type authService struct {
  userRepo repositories.UserRepository
}

// NewAuthService - constructor
func NewAuthService(userRepo repositories.UserRepository) AuthService {
  return &authService{
    userRepo: userRepo,
  }
}

// Login - handle logic for login
func (s *authService) Login(ctx context.Context, input types.LoginInput) (*types.LoginResponse, error) {
  // Set timeout for query
  ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
  defer cancel()

  // Find user by email
  user, err := s.userRepo.FindByEmail(ctx, input.Email)
  if err != nil {
    return nil, errors.New("invalid credentials")
  }

  // Verify password
  if !utils.CheckPassword(user.Password, input.Password) {
    return nil, errors.New("invalid credentials")
  }

  // Generate JWT token
  token, err := utils.GenerateToken(user.ID, user.Email)
  if err != nil {
    return nil, errors.New("failed to generate token")
  }

  // Convert to response DTO
  userResponse := types.ToUserResponse(user)

  return &types.LoginResponse{
    Token: token,
    User:  userResponse,
  }, nil
}
