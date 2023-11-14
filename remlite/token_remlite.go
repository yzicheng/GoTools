package remlite

import (
	"context"
	"time"
)

type LimitHandel func(ctx context.Context)

type TokenLimit struct {
	token chan struct{}
}

func NewBuildTokenLimit(duration time.Duration, size int) *TokenLimit {
	limit := &TokenLimit{token: make(chan struct{}, size)}
	go func() {
		tick := time.NewTicker(duration)
		defer tick.Stop()
		for {
			select {
			case <-tick.C:
				select {
				// 如果桶没有满就塞入,否则进入default
				case limit.token <- struct{}{}:
				default:
					// 如果桶是满的就过去
				}
			}
		}
	}()
	return limit
}

func (l *TokenLimit) BuildServer() {

}
