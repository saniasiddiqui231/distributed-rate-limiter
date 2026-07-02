package limiter

import (
	_ "embed"
	"time"

	goredis "github.com/redis/go-redis/v9"

	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/config"
	redisClient "github.com/saniasiddiqui231/distributed-rate-limiter/internal/redis"
)

//go:embed token_limiter.lua
var luaScript string

type TokenBucket struct {
	script *goredis.Script
}

func NewTokenBucket() *TokenBucket {
	return &TokenBucket{
		script: goredis.NewScript(luaScript),
	}
}

func (l *TokenBucket) Allow(clientID string) (bool, error) {

	key := "rate_limit:" + clientID

	now := time.Now().Unix()

	result, err := l.script.Run(
		redisClient.Ctx,
		redisClient.Client,
		[]string{key},

		config.AppConfig.BucketCapacity,
		config.AppConfig.RefillTokens,
		int(config.AppConfig.RefillInterval.Seconds()),
		now,
		int(config.AppConfig.KeyTTL.Seconds()),
	).Int()

	if err != nil {
		return false, err
	}

	return result == 1, nil
}
