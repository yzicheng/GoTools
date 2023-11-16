package remlite

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTokenLimit_BuildHttpServerInterceptor(t1 *testing.T) {
	tests := []struct {
		name      string
		ctx       context.Context
		b         func() TokenLimitHttpHandel
		handel    func(ctx context.Context, req interface{}) (resp interface{}, err error)
		req       interface{}
		wantError error
		wantResp  interface{}
	}{
		{
			name: "close",
			ctx:  context.Background(),
			req:  "",
			b: func() TokenLimitHttpHandel {
				t := NewBuildTokenLimit(time.Second, 1)
				interceptor := t.BuildHttpServerInterceptor()
				defer t.Close(context.Background())
				return interceptor
			},
			handel: func(ctx context.Context, req interface{}) (resp interface{}, err error) {
				return nil, nil
			},
			wantError: errors.New("令牌已被关闭"),
			wantResp:  nil,
		},
		{
			name: "ctx cancel",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				return ctx
			}(),
			req: "",
			b: func() TokenLimitHttpHandel {
				t := NewBuildTokenLimit(time.Second, 1)
				interceptor := t.BuildHttpServerInterceptor()
				return interceptor
			},
			handel: func(ctx context.Context, req interface{}) (resp interface{}, err error) {
				return nil, nil
			},
			wantError: errors.New("等待超时"),
			wantResp:  nil,
		},
		{
			name: "normal",
			ctx: func() context.Context {
				return context.Background()
			}(),
			req: "",
			b: func() TokenLimitHttpHandel {
				t := NewBuildTokenLimit(time.Millisecond, 1)
				interceptor := t.BuildHttpServerInterceptor()
				return interceptor
			},
			handel: func(ctx context.Context, req interface{}) (resp interface{}, err error) {
				return "hello world", nil
			},
			wantError: nil,
			wantResp:  "hello world",
		},
		{
			name: "limit",
			ctx: func() context.Context {
				return context.Background()
			}(),
			req: "",
			b: func() TokenLimitHttpHandel {
				t := NewBuildTokenLimit(time.Second*2, 1)
				interceptor := t.BuildHttpServerInterceptor()
				return interceptor
			},
			handel: func(ctx context.Context, req interface{}) (resp interface{}, err error) {
				return nil, nil
			},
			wantError: errors.New("到达瓶颈"),
			wantResp:  nil,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			interceptor := tt.b()
			time.Sleep(time.Second)
			res, err := interceptor(tt.ctx, tt.req, tt.handel)
			assert.Equal(t1, err, tt.wantError)
			if err != nil {
				return
			}
			assert.Equal(t1, res, tt.wantResp)

		})
	}
}
