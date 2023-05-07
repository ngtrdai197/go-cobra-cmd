package redis

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	MAX_RETRY_BACKOFF = 500 * time.Millisecond
	MAX_RETRIES       = 5
	WRITE_TIMEOUT     = 3 * time.Second
)

type Adaptor struct {
	Redis *redis.Client
}

// Pass redisURL with format localhost:6379/0
func NewRedisConnection(redisURL string) (*Adaptor, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("Can't connection redis : %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:            opts.Addr,
		DB:              opts.DB,
		MaxRetries:      MAX_RETRIES,
		MaxRetryBackoff: MAX_RETRY_BACKOFF,
		WriteTimeout:    WRITE_TIMEOUT,
	})

	return &Adaptor{
		Redis: redisClient,
	}, nil
}
