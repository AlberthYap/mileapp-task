package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func LoggerMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    start := time.Now()
    path := c.Request.URL.Path
    method := c.Request.Method

    // Process request
    c.Next()

    // Log after request
    latency := time.Since(start)
    statusCode := c.Writer.Status()

    log.Info().
      Str("method", method).
      Str("path", path).
      Int("status", statusCode).
      Dur("latency", latency).
      Str("ip", c.ClientIP()).
      Str("user_agent", c.GetHeader("User-Agent")).
      Msg("HTTP request")
  }
}
