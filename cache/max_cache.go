package cache

import (
	"context"
	"errors"
	"sync/atomic"
	"time"
)

var (
	errorOutOfRange = errors.New("key超过最大内存")
)

type MaxCache struct {
	*LocalCache
	cnt    int32
	maxCnt int32
}

func BuildMaxCache(cache *LocalCache, maxCnt int32) *MaxCache {
	maxCache := &MaxCache{
		cache,
		0,
		maxCnt,
	}
	evict := cache.evict
	maxCache.evict = func(key string, val any) {
		atomic.AddInt32(&maxCache.cnt, -1)
		evict(key, val)
	}
	return maxCache
}

func (maxCache *MaxCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	maxCache.Lock()
	defer maxCache.Unlock()
	_, ok := maxCache.cache[key]
	if !ok {
		if maxCache.cnt+1 > maxCache.maxCnt {
			// 实现淘汰策略
			return errorOutOfRange
		}
		maxCache.cnt++
	}
	return maxCache.set(ctx, key, val, expiration)
}
