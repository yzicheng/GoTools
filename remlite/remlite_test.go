package remlite

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

// Go 官方的一个限流器简单应用
// 他本身基于锁实现,思考是否可以做一个无锁设计
func TestRemlite(t *testing.T) {
	// golang.org/x/time/rate 是一个有锁的操作
	lt := rate.NewLimiter(rate.Every(time.Second/10), 10)
	go func() {
		for i := 0; i < 1000; i++ {
			err := lt.WaitN(context.Background(), 1)
			if err != nil {
				return
			}
			fmt.Println(i)
		}
	}()
	time.Sleep(time.Hour)
}
