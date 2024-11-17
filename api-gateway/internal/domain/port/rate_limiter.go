package port

import "context"

type ReteLimiterRepository interface {
	Allow(ctx context.Context, key string) bool
}
