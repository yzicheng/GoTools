package cache

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	errorSetFail = errors.New("redis插入失败")
)

type RedisCache struct {
	client redis.Cmdable
}

func NewRedisCache(client redis.Cmdable) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

func (r RedisCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	set, err := r.client.Set(ctx, key, val, expiration).Result()
	if err != nil {
		return err
	}
	if set == "OK" {
		return nil
	}
	return fmt.Errorf("%s,%s", errorSetFail, key)
}

func (r RedisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r RedisCache) Delete(ctx context.Context, key string) error {
	_, err := r.client.Del(ctx, key).Result()
	return err
}
