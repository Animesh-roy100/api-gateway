package cache

import (
	"api-gateway/internal/domain/port"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr string) port.CacheRepository {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &RedisCache{
		client: client,
	}
}

func (r *RedisCache) Get(ctx context.Context, key string) ([]byte, bool) {
	val, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, false
	}

	return val, true
}

func (r *RedisCache) Set(ctx context.Context, key string, value []byte, expiration time.Duration) {
	r.client.Set(ctx, key, value, expiration)
}
