package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"task-api/configs"
)

// SetupTrustedProxies configures trusted proxies based on TRUSTED_PROXIES env or GIN_MODE
func SetupTrustedProxies(r *gin.Engine) {
	var proxies []string

	// Priority 1: Check custom TRUSTED_PROXIES env variable
	trustedProxies := configs.GetEnv("TRUSTED_PROXIES", "")
	if trustedProxies != "" {
		if trustedProxies == "none" {
			proxies = nil
			log.Info().Msg("Trusted proxies explicitly disabled via TRUSTED_PROXIES=none")
		} else {
			proxies = strings.Split(trustedProxies, ",")
			for i := range proxies {
				proxies[i] = strings.TrimSpace(proxies[i])
			}

			log.Info().Strs("proxies", proxies).Msg("Using custom trusted proxies from TRUSTED_PROXIES")
		}
	} else {
		// Priority 2: Use defaults based on GIN_MODE
		ginMode := gin.Mode()

		switch ginMode {
		case gin.ReleaseMode:
			proxies = []string{
				"10.0.0.0/8",
				"172.16.0.0/12",
				"192.168.0.0/16",
			}
			log.Info().Str("mode", ginMode).Msg("Using production default trusted proxies")

		case gin.DebugMode, gin.TestMode:
			proxies = nil
			log.Info().Str("mode", ginMode).Msg("No trusted proxies (development mode)")
		}
	}

	if err := r.SetTrustedProxies(proxies); err != nil {
		log.Fatal().Err(err).Msg("Failed to set trusted proxies")
	}

	if len(proxies) > 0 {
		log.Info().Strs("configured_proxies", proxies).Msg("Trusted proxies configured successfully")
	} else {
		log.Info().Msg("No trusted proxies configured")
	}
}
