package limiter

import (
	"fmt"

	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/config"
)

func New() (Limiter, error) {

	switch config.AppConfig.Algorithm {

	case "token_bucket":
		return NewTokenBucket(), nil

	case "sliding_window":
		return NewSlidingWindow(), nil

	default:
		return nil, fmt.Errorf("unknown algorithm: %s", config.AppConfig.Algorithm)
	}
}
