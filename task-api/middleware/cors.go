package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"task-api/configs"
)

func CORSMiddleware() gin.HandlerFunc {
    // Get allowed origins from env
    allowedOriginsStr := configs.GetEnv("ALLOWED_ORIGINS", "http://localhost:5173")
    allowedOrigins := strings.Split(allowedOriginsStr, ",")
    
    // Trim whitespace
    for i := range allowedOrigins {
        allowedOrigins[i] = strings.TrimSpace(allowedOrigins[i])
    }

    log.Info().Strs("allowed_origins", allowedOrigins).Msg("CORS configured")

    return func(c *gin.Context) {
        origin := c.Request.Header.Get("Origin")

        // Check if origin is in allowed list
        isAllowed := false
        for _, allowedOrigin := range allowedOrigins {
            if allowedOrigin == "*" || origin == allowedOrigin {
                isAllowed = true
                break
            }
        }

        if isAllowed {
            if origin != "" {
                c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
            } else {
                c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigins[0])
            }
            c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        }

        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
