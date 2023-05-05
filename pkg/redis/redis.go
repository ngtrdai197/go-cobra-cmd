package redis

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	maxRetryBackoff = 500 * time.Millisecond
	maxRetries      = 5
	writeTimeout    = 3 * time.Second
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
		MaxRetries:      maxRetries,
		MaxRetryBackoff: maxRetryBackoff,
		WriteTimeout:    writeTimeout,
	})

	return &Adaptor{
		Redis: redisClient,
	}, nil

}
