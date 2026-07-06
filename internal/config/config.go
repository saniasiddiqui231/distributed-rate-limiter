package config

import (
	"os"
	"time"
)

type Config struct {
	Algorithm string

	// Shared rate limit policy
	RateLimit int
	Window    time.Duration

	// Token bucket specific
	RefillTokens   float64
	RefillInterval time.Duration

	KeyTTL time.Duration

	RedisAddr  string
	BackendURL string
	ServerPort string
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

var AppConfig = Config{
	Algorithm: "token_bucket", // change to "sliding_window" to switch

	RateLimit: 5,
	Window:    10 * time.Second,

	RefillTokens:   1,
	RefillInterval: time.Second,

	KeyTTL: 2 * time.Minute,

	RedisAddr:  getEnv("REDIS_ADDR", "localhost:6379"),
	BackendURL: getEnv("BACKEND_URL", "http://localhost:8081"),
	ServerPort: getEnv("SERVER_PORT", ":8080"),
}
