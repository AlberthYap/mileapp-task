package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"task-api/utils"
)

// AuthMiddleware - JWT authentication
func AuthMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    // Get token from header
    authHeader := c.GetHeader("Authorization")
    
    if authHeader == "" {
      log.Warn().Str("ip", c.ClientIP()).Msg("Missing authorization header")
      utils.Fail(c, 401, "Unauthorized", gin.H{"error": "Authorization header required"})
      c.Abort()
      return
    }
    
    // Remove "Bearer " prefix
    token := authHeader
    if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
      token = authHeader[7:]
    }
    
    // Validate token
    claims, err := utils.ValidateToken(token)
    if err != nil {
      log.Warn().Err(err).Str("ip", c.ClientIP()).Msg("Invalid token")
      utils.Fail(c, 401, "Unauthorized", gin.H{"error": "Invalid or expired token"})
      c.Abort()
      return
    }
    
    // Set user info in context
    c.Set("userID", claims.UserID)
    c.Set("userEmail", claims.Email)
    
    c.Next()
  }
}
