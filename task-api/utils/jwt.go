package utils

import (
	"errors"
	"os"
	"task-api/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GenerateToken - generate JWT token
func GenerateToken(userID bson.ObjectID, email string) (string, error) {
  jwtSecret := os.Getenv("JWT_SECRET")
  if jwtSecret == "" {
    return "", errors.New("JWT_SECRET not configured")
  }

	// set token data
  claims := types.JWTClaims{
    UserID: userID,
    Email:  email,
    RegisteredClaims: jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
      IssuedAt:  jwt.NewNumericDate(time.Now()),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString([]byte(jwtSecret))
}

// ValidateToken - validate JWT token
func ValidateToken(tokenString string) (*types.JWTClaims, error) {
  jwtSecret := os.Getenv("JWT_SECRET")
  if jwtSecret == "" {
    return nil, errors.New("JWT_SECRET not configured")
  }

  token, err := jwt.ParseWithClaims(tokenString, &types.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
    return []byte(jwtSecret), nil
  })

  if err != nil {
    return nil, err
  }

  claims, ok := token.Claims.(*types.JWTClaims)
  if !ok || !token.Valid {
    return nil, errors.New("invalid token")
  }

  return claims, nil
}
