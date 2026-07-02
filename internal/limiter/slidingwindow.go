package limiter

import (
	"time"

	"github.com/saniasiddiqui231/distributed-rate-limiter/internal/redis"
)

const (
	MaxRequests = 5
	Window      = 10 * time.Second
)

func Allow(clientID string) (bool, error) {

	count, err := redis.Client.Incr(redis.Ctx, clientID).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		err = redis.Client.Expire(redis.Ctx, clientID, Window).Err()
		if err != nil {
			return false, err
		}
	}

	if count > MaxRequests {
		return false, nil
	}

	return true, nil
}
