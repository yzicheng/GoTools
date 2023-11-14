package cache

import (
	"Tool/cache/mocks"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	testCase := []struct {
		mock       func(ctrl *gomock.Controller) redis.Cmdable
		name       string
		key        string
		val        string
		expiration time.Duration
		want       string
		wantErr    error
	}{
		{
			name: "kye1",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmdable := mocks.NewMockCmdable(ctrl)
				status := redis.NewStatusCmd(context.Background())
				status.SetVal("OK")
				cmdable.EXPECT().
					Set(context.Background(), "key1", "value1", time.Second).Return(status)
				return cmdable
			},
			key:        "key1",
			val:        "value1",
			expiration: time.Second,
			wantErr:    nil,
		},
		{
			name: "expiration",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmdable := mocks.NewMockCmdable(ctrl)
				status := redis.NewStatusCmd(context.Background())
				status.SetErr(context.DeadlineExceeded)
				cmdable.EXPECT().
					Set(context.Background(), "key1", "value1", time.Second).Return(status)
				return cmdable
			},
			key:        "key1",
			val:        "value1",
			expiration: time.Second,
			wantErr:    context.DeadlineExceeded,
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := NewRedisCache(tt.mock(ctrl))
			err := client.Set(context.Background(), tt.key, tt.val, tt.expiration)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGet(t *testing.T) {
	testCase := []struct {
		mock       func(ctrl *gomock.Controller) redis.Cmdable
		name       string
		key        string
		val        string
		expiration time.Duration
		want       string
		wantErr    error
	}{
		{
			name: "kye1",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmdable := mocks.NewMockCmdable(ctrl)
				status := redis.NewStatusCmd(context.Background())
				status.SetVal("OK")
				cmdable.EXPECT().
					Set(context.Background(), "key1", "value1", time.Second).Return(status)
				return cmdable
			},
			key:        "key1",
			val:        "value1",
			expiration: time.Second,
			wantErr:    nil,
		},
		{
			name: "expiration",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmdable := mocks.NewMockCmdable(ctrl)
				status := redis.NewStatusCmd(context.Background())
				status.SetErr(context.DeadlineExceeded)
				cmdable.EXPECT().
					Set(context.Background(), "key1", "value1", time.Second).Return(status)
				return cmdable
			},
			key:        "key1",
			val:        "value1",
			expiration: time.Second,
			wantErr:    context.DeadlineExceeded,
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := NewRedisCache(tt.mock(ctrl))
			err := client.Set(context.Background(), tt.key, tt.val, tt.expiration)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestPro(t *testing.T) {
	//fmt.Println(test1())
	//fmt.Println(test2())
	//fmt.Println(test3())
	fmt.Println(test4())
	return
}

func test1() (v int) {
	defer fmt.Println("test1")
	return v
}

func test2() (v int) {
	defer func() {
		fmt.Println("test2")
	}()
	return 3
}

func test3() (v int) {
	defer fmt.Println(v)
	v = 3
	return 4
}

func test4() (v int) {
	defer func(n int) {
		fmt.Println(n)
	}(v)
	return 5
}
