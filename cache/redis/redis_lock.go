package redis

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

var lock, refresh, unlock string

func init() {
	// 读区本地lua文件
	lua, err := os.ReadFile("./lua/lock.lua")
	if err != nil {
		panic(err)
	}
	lock = string(lua)

	lua, err = os.ReadFile("./lua/refresh.lua")
	if err != nil {
		panic(err)
	}
	refresh = string(lua)

	lua, err = os.ReadFile("./lua/unlock.lua")
	if err != nil {
		panic(err)
	}
	unlock = string(lua)
}

type RedisMutex interface {
	Lock(string) (string, bool)
	UnLock(string, string) bool
	Refresh(string, string, time.Duration) bool
}

type RedisMuetxClient struct {
	ctx        context.Context
	minTimeout time.Duration // 最小超时时间
	minExpire  time.Duration // 最小超过期时间
	minReExec  time.Duration // 重试间隔
	client     redis.Client
}

func NewRedisMuetxClient(ctx context.Context, client redis.Client) *RedisMuetxClient {
	return &RedisMuetxClient{
		ctx:        ctx,
		minTimeout: 30 * time.Millisecond,
		minExpire:  5 * time.Millisecond,
		minReExec:  1 * time.Millisecond,
		client:     client,
	}
}

func (r RedisMuetxClient) Lock(key string, timeout time.Duration, expire time.Duration) (string, bool) {
	if timeout < r.minTimeout {
		timeout = r.minTimeout
	}
	if expire < r.minExpire {
		expire = r.minExpire
	}
	// 重试间隔
	token := uuid.New().String()
	reExec := int(timeout / r.minReExec)
	for i := 0; i < reExec; i++ {
		// 加锁
		script := redis.NewScript(lock)
		result, err := script.Run(context.Background(), r.client, []string{key}, token, expire).Result()
		if err != nil {
			fmt.Println("执行脚本错误：", err)
			return "", false
		}
		if result == nil || result == "" {
			fmt.Println("锁被人拿着或者出现其他错误")
			time.Sleep(r.minReExec)
		} else {
			return token, true
		}
	}
	return "", false
}

func (r RedisMuetxClient) UnLock(key string, token string) bool {
	// 解锁
	script := redis.NewScript(unlock)
	result, err := script.Run(context.Background(), r.client, []string{key}, token).Result()
	if err != nil {
		fmt.Println("执行脚本错误：", err)
		return false
	}
	if result == 0 {
		return false
	}
	return true
}
func (r RedisMuetxClient) Refresh(key string, token string, expire time.Duration) bool {
	// 续期
	script := redis.NewScript(refresh)
	result, err := script.Run(context.Background(), r.client, []string{key}, token, expire).Result()
	if err != nil {
		fmt.Println("执行脚本错误：", err)
		return false
	}
	if result == 0 {
		return false
	}
	return true
}
