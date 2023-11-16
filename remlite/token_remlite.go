package remlite

import (
	"Tool/pool"
	"context"
	"errors"
	"google.golang.org/grpc"
	"io"
	"net/http"
	"time"
)

type TokenLimitGet func(ctx context.Context, url string, connection pool.Connection) (*http.Response, error)
type TokenLimitPost func(ctx context.Context, url, contentType string, body io.Reader, connection pool.Connection) (resp *http.Response, err error)
type TokenLimitHttpHandel func(ctx context.Context, req interface{}, handel func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error)

type TokenLimit struct {
	token chan struct{}
	close chan struct{}
}

func NewBuildTokenLimit(duration time.Duration, size int) *TokenLimit {
	ch := make(chan struct{}, size)
	closeCh := make(chan struct{})
	tick := time.NewTicker(duration)
	go func() {
		defer tick.Stop()
		for {
			select {
			case <-tick.C:
				select {
				// 如果桶没有满就塞入,否则进入default
				case ch <- struct{}{}:
				default:
				}
			// 主动关闭
			case <-closeCh:
				return
			}
		}
	}()
	return &TokenLimit{token: ch, close: closeCh}
}

// BuildServerInterceptor Grpc服务端限流
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

// BuildClientInterceptor Grpc客户端限流
func (t *TokenLimit) BuildClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// 要在这里拿到令牌
		select {
		case <-t.close:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		case <-t.token:
			err := invoker(ctx, method, req, reply, cc, opts...)
			if err != nil {
				return err
			}
		default:
			return errors.New("到达瓶颈")

		}
		return nil
	}

}

// BuildHttpClientGetInterceptor http客户端Get端限流
func (t *TokenLimit) BuildHttpClientGetInterceptor() TokenLimitGet {
	return func(ctx context.Context, url string, connection pool.Connection) (*http.Response, error) {
		// 要在这里拿到令牌
		select {
		case <-t.close:
			return nil, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-t.token:
			return connection.Get(ctx, url)
		default:
			return nil, errors.New("到达瓶颈")

		}
	}
}

// BuildHttpClientPostInterceptor http客户端Post端限流
func (t *TokenLimit) BuildHttpClientPostInterceptor() TokenLimitPost {
	return func(ctx context.Context, url, contentType string, body io.Reader, connection pool.Connection) (resp *http.Response, err error) {
		// 要在这里拿到令牌
		select {
		case <-t.close:
			return nil, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-t.token:
			return connection.Post(ctx, url, contentType, body)
		default:
			return nil, errors.New("到达瓶颈")
		}
	}
}

// BuildHttpServerInterceptor http服务端限流
func (t *TokenLimit) BuildHttpServerInterceptor() TokenLimitHttpHandel {
	return func(ctx context.Context, req interface{}, handel func(ctx context.Context, req interface{}) (resp interface{}, err error)) (resp interface{}, err error) {
		// 要在这里拿到令牌
		select {
		case <-t.close:
			return nil, errors.New("令牌已被关闭")
		case <-ctx.Done():
			return nil, errors.New("等待超时")
		case <-t.token:
			return handel(ctx, req)
		default:
			return nil, errors.New("到达瓶颈")
		}
	}
}

func (t *TokenLimit) Close(ctx context.Context) error {
	// 不能用<-t.close,我有两个地方监听了close
	//<-t.close
	close(t.close)
	return nil
}
