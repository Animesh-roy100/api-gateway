package port

import (
	"context"
	"time"
)

type CacheRepository interface {
	Get(ctx context.Context, key string) ([]byte, bool)
	Set(ctx context.Context, key string, value []byte, expiration time.Duration)
}
