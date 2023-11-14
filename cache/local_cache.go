package cache

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	errorKeyIsNotVisit = errors.New("当前key不存在")
	errorCloseRep      = errors.New("重复关闭")
)

type LocalCache struct {
	cache map[string]*item
	close chan struct{} // 控制可以关闭监听的gorouting
	evict func(key string, val any)
	sync.RWMutex
}

type BuildLocalCacheOption func(cache *LocalCache)

func BuildLocalCacheWithEvict(fn func(key string, val any)) BuildLocalCacheOption {
	return func(cache *LocalCache) {
		cache.evict = fn
	}
}

// NewBuildLocalCache 创建一个本地缓存队列
// 开启定时任务删除过期key
// 删除条件:每间expiration.隔轮询1000次.
func NewBuildLocalCache(size int, expiration time.Duration, opts ...BuildLocalCacheOption) *LocalCache {
	res := &LocalCache{
		cache: make(map[string]*item, size),
		close: make(chan struct{}),
		evict: func(key string, val any) {

		},
	}
	for _, opt := range opts {
		opt(res)
	}
	tick := time.NewTicker(expiration)
	go func() {
		// 优化方案
		// 这里可以使用延迟队列或者维持一个树形结构,过期的数据都在队首或者在叶子节点
		for {
			select {
			case t := <-tick.C:
				times := 0
				// 如果这里锁住的内容发生异常,需要业务去调用defer Close
				res.Lock()
				// 用了map随机寻值的特性解决了始终从第一个遍历的问题
				for k, v := range res.cache {
					if times > 1000 {
						break
					}
					if v.deadLineBefore(t) {
						res.delete(k)
						//delete(res.cache, k)
					}
					times++
				}
				res.Unlock()
			case <-res.close:
				return
			}
		}
	}()
	return res
}

// Set 缓存中增加一个键值对
// 若expiration为0时,永久保存
// 增加了写锁,所以是线程安全的
func (l *LocalCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	l.Lock()
	defer l.Unlock()
	l.set(ctx, key, val, expiration)
	// 这里会有一种情况,当A设置了Key的Value 过期时间1分钟
	// B在期间又覆盖了Key那么在A到期后就会清除B
	//if expiration > 0 {
	//	time.AfterFunc(expiration, func() {
	//		l.Lock()
	//		defer l.Unlock()
	//		if v, ok := l.cache[key]; ok && v.deadline.Before(time.Now()) {
	//			delete(l.cache, key) // 这里删除掉key
	//		}
	//	})
	//}
	return nil
}
func (l *LocalCache) set(ctx context.Context, key string, val any, expiration time.Duration) error {
	l.cache[key] = &item{
		val: val,
	}
	dl := time.Now().Add(expiration)
	if expiration > 0 {
		l.cache[key].deadline = dl
	}
	return nil
}

// Get 从缓存中获取一个键值队伍
// 参考redis删除缓存策略,在获取key时检查是否过期
// 若过期则删除
// 删除过程增加写锁,所以是线程安全的,并且在写锁前后都增加了校验（double check）所以不会出现错删
func (l *LocalCache) Get(ctx context.Context, key string) (any, error) {
	l.RLock()
	v, ok := l.cache[key]
	l.RUnlock()
	if !ok {
		return nil, fmt.Errorf("%w ,%s", errorKeyIsNotVisit, key)
	} else {
		// double-check
		// 第一遍检查
		if v.deadLineBefore(time.Now()) {
			l.Lock()
			defer l.Unlock()
			// 防止加锁前又被更改了 再次检查一遍
			v, ok = l.cache[key]
			if !ok {
				return nil, fmt.Errorf("%w ,%s", errorKeyIsNotVisit, key)
			}
			if v.deadLineBefore(time.Now()) {
				l.delete(key)
			}
		}
		return v.val, nil
	}
}

// Delete 删除一个本地缓存
// 增加了写锁,所以是线程安全的
func (l *LocalCache) Delete(ctx context.Context, key string) error {
	l.Lock()
	defer l.Unlock()
	l.delete(key)
	return nil
}

func (l *LocalCache) delete(key string) {
	v, ok := l.cache[key]
	if !ok {
		return
	}
	delete(l.cache, key)
	l.evict(key, v)
}

// Close 关闭本地缓存的结构体,停止定时任务检查过期key
func (l *LocalCache) Close() error {
	select {
	case l.close <- struct{}{}:
	default:
		return fmt.Errorf("%s", errorCloseRep)
	}
	return nil
}

type item struct {
	val      any
	deadline time.Time
}

func (itm *item) deadLineBefore(t time.Time) bool {
	return !itm.deadline.IsZero() && itm.deadline.Before(t)
}
