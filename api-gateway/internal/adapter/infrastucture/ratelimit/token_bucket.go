package ratelimit

import (
	"api-gateway/internal/domain/port"
	"context"
	"sync"

	"golang.org/x/time/rate"
)

type TokenBucketLimiter struct {
	limiters sync.Map
	rate     rate.Limit
	burst    int
}

func NewTokenBucketLimiter(r rate.Limit, b int) port.ReteLimiterRepository {
	return &TokenBucketLimiter{
		rate:  r,
		burst: b,
	}
}

func (t *TokenBucketLimiter) Allow(ctx context.Context, key string) bool {
	limiter, _ := t.limiters.LoadOrStore(key, rate.NewLimiter(t.rate, t.burst))
	return limiter.(*rate.Limiter).Allow()
}
