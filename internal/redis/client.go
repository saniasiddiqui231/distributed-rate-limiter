package redis

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/config"
)

var Ctx = context.Background()

var Client = redis.NewClient(&redis.Options{
	Addr: config.AppConfig.RedisAddr,
})
