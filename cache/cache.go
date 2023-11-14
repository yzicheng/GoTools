package cache

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, val any, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}
