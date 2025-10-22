// utils/logger.go
package utils

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
  // Create logs directory if not exists
  if err := os.MkdirAll("./logs", 0755); err != nil {
    panic(err)
  }

  // Open log file
  logFile, err := os.OpenFile(
    "./logs/app.log",
    os.O_CREATE|os.O_WRONLY|os.O_APPEND,
    0666,
  )
  if err != nil {
    panic(err)
  }

  if os.Getenv("GIN_MODE") == "release" {
    // Production: Write JSON to file
    log.Logger = zerolog.New(logFile).With().Timestamp().Logger()
  } else {
    // Development: Pretty console + file
    multi := zerolog.MultiLevelWriter(
      zerolog.ConsoleWriter{Out: os.Stdout},  // Console (pretty)
      logFile,                                 // File (JSON)
    )
    log.Logger = zerolog.New(multi).With().Timestamp().Logger()
  }

  log.Info().Msg("Logger initialized")
}
