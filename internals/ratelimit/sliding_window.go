package ratelimit

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type SlidingWindow struct {
	client *redis.Client
	limit  int
	window time.Duration
}

func NewSlidingWindow(client *redis.Client, limit int, window time.Duration) *SlidingWindow {
	return &SlidingWindow{
		client: client,
		limit:  limit,
		window: window,
	}
}

func (s *SlidingWindow) Allow(ctx context.Context, key string) (bool, int, error) {
	return true, 1, nil
}
