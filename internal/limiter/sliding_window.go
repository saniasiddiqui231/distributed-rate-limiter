package limiter

import (
	_ "embed"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"

	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/config"
	redisClient "github.com/saniasiddiqui231/distributed-rate-limiter/internal/redis"
)

//go:embed slidingwindow.lua
var slidingWindowLua string

type SlidingWindow struct {
	script *goredis.Script
}

func NewSlidingWindow() *SlidingWindow {
	return &SlidingWindow{
		script: goredis.NewScript(slidingWindowLua),
	}
}

func (s *SlidingWindow) Allow(clientID string) (bool, error) {

	key := "rate_limit:" + clientID

	now := time.Now()

	timestamp := now.Unix()

	member := fmt.Sprintf(
		"%d-%d",
		timestamp,
		now.UnixNano(),
	)

	result, err := s.script.Run(
		redisClient.Ctx,
		redisClient.Client,
		[]string{key},
		config.AppConfig.RateLimit,
		int(config.AppConfig.Window.Seconds()),
		timestamp,
		member,
	).Int()

	if err != nil {
		return false, err
	}

	return result == 1, nil
}
