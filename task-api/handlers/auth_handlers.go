package handlers

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"task-api/services"
	"task-api/types"
	"task-api/utils"
)

// AuthHandler - struct for auth handlers
type AuthHandler struct {
  authService services.AuthService
}

// NewAuthHandler - constructor
func NewAuthHandler(authService services.AuthService) *AuthHandler {
  return &AuthHandler{
    authService: authService,
  }
}

// Login - handler for login
func (h *AuthHandler) Login(c *gin.Context) {
  var input types.LoginInput

  // Handle empty body (EOF error)
  if err := c.ShouldBindJSON(&input); err != nil {
    if err == io.EOF {
      utils.Fail(c, http.StatusBadRequest, "Request body required", gin.H{
        "error": "Please provide email and password",
      })
      return
    }

    utils.Fail(c, http.StatusBadRequest, types.MsgValidationFailed, gin.H{
      "error": err.Error(),
    })
    return
  }

  // Add context timeout
  ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
  defer cancel()

  // Add logging
  log.Printf("Login attempt: email=%s, ip=%s", input.Email, c.ClientIP())

  // Call service
  response, err := h.authService.Login(ctx, input)
  if err != nil {
    // Log failed attempts
    log.Printf("Login failed: email=%s, error=%s", input.Email, err.Error())
    
    utils.Fail(c, http.StatusUnauthorized, types.MsgLoginFailed, gin.H{
      "error": types.MsgInvalidCredentials,
    })
    return
  }

  // Log successful login
  log.Printf("Login successful: email=%s", input.Email)

  // Success response
  utils.Success(c, http.StatusOK, types.MsgLoginSuccess, response)
}