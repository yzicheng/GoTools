package remlite

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

type LimitHandel func(ctx context.Context, req interface{}) interface{}

type TokenLimit struct {
	token chan struct{}
	close chan struct{}
}

func NewBuildTokenLimit(duration time.Duration, size int) *TokenLimit {
	ch := make(chan struct{}, size)
	close := make(chan struct{})
	tick := time.NewTicker(duration)
	go func() {
		defer tick.Stop()
		for {
			select {
			case <-tick.C:
				select {
				// 主动关闭
				case <-close:
					return
				// 如果桶没有满就塞入,否则进入default
				case ch <- struct{}{}:
				default:
					// 如果桶是满的就过去
					fmt.Println("桶满了")
				}
			}
		}
	}()
	return &TokenLimit{token: ch, close: close}
}

func (t *TokenLimit) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 要在这里拿到令牌
		select {
		case <-t.close:
			return
		case <-ctx.Done():
			err = ctx.Err()
			return
		case <-t.token:
			resp, err = handler(ctx, req)
		default:
			err = errors.New("到达瓶颈")
			return
		}
		return
	}
}

func (t *TokenLimit) Close(ctx context.Context) error {
	// 不能用<-t.close,我有两个地方监听了close
	//<-t.close
	close(t.close)
	return nil
}
