package cache

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestBuildLocalCacheWithEvict(t *testing.T) {
	type args struct {
		fn func(key string, val any)
	}
	tests := []struct {
		name string
		args args
		want BuildLocalCacheOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildLocalCacheWithEvict(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildLocalCacheWithEvict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalCache_Close(t *testing.T) {
	type fields struct {
		cache   map[string]*item
		close   chan struct{}
		evict   func(key string, val any)
		RWMutex sync.RWMutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LocalCache{
				cache:   tt.fields.cache,
				close:   tt.fields.close,
				evict:   tt.fields.evict,
				RWMutex: tt.fields.RWMutex,
			}
			if err := l.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalCache_Delete(t *testing.T) {
	type fields struct {
		cache   map[string]*item
		close   chan struct{}
		evict   func(key string, val any)
		RWMutex sync.RWMutex
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LocalCache{
				cache:   tt.fields.cache,
				close:   tt.fields.close,
				evict:   tt.fields.evict,
				RWMutex: tt.fields.RWMutex,
			}
			if err := l.Delete(tt.args.ctx, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalCache_Get(t *testing.T) {
	tests := []struct {
		name    string
		cache   func() *LocalCache
		input   string
		want    any
		wantErr error
	}{
		{
			name: "key is not fount",
			cache: func() *LocalCache {
				return NewBuildLocalCache(100, 3*time.Second)
			},
			input:   "key is not fount",
			wantErr: fmt.Errorf("%w ,%s", errorKeyIsNotVisit, "key is not fount"),
		},
		{
			name: "key1",
			cache: func() *LocalCache {
				localCache := NewBuildLocalCache(100, 3*time.Second)
				err := localCache.Set(context.Background(), "key1", "key1", 0)
				require.NoError(t, err)
				return localCache
			},
			input:   "key1",
			wantErr: nil,
			want:    "key1",
		},
		{
			name: "expiration",
			cache: func() *LocalCache {
				localCache := NewBuildLocalCache(100, time.Second)
				err := localCache.Set(context.Background(), "key1", "key1", time.Second)
				require.NoError(t, err)
				time.Sleep(3 * time.Second)
				return localCache
			},
			input:   "key1",
			wantErr: fmt.Errorf("%w ,%s", errorKeyIsNotVisit, "key1"),
			want:    "key1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get, err := tt.cache().Get(context.Background(), tt.input)
			if err != nil {
				assert.Error(t, err, tt.wantErr)
				return
			}
			assert.Equal(t, get, tt.want)
		})
	}
}

func TestLocalCache_Set(t *testing.T) {
	type fields struct {
		cache   map[string]*item
		close   chan struct{}
		evict   func(key string, val any)
		RWMutex sync.RWMutex
	}
	type args struct {
		ctx        context.Context
		key        string
		val        any
		expiration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LocalCache{
				cache:   tt.fields.cache,
				close:   tt.fields.close,
				evict:   tt.fields.evict,
				RWMutex: tt.fields.RWMutex,
			}
			if err := l.Set(tt.args.ctx, tt.args.key, tt.args.val, tt.args.expiration); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalCache_delete(t *testing.T) {
	type fields struct {
		cache   map[string]*item
		close   chan struct{}
		evict   func(key string, val any)
		RWMutex sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LocalCache{
				cache:   tt.fields.cache,
				close:   tt.fields.close,
				evict:   tt.fields.evict,
				RWMutex: tt.fields.RWMutex,
			}
			l.delete(tt.args.key)
		})
	}
}

func TestNewBuildLocalCache(t *testing.T) {
	type args struct {
		size       int
		expiration time.Duration
		opts       []BuildLocalCacheOption
	}
	tests := []struct {
		name string
		args args
		want *LocalCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuildLocalCache(tt.args.size, tt.args.expiration, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuildLocalCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_item_deadLineBefore(t *testing.T) {
	type fields struct {
		val      any
		deadline time.Time
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itm := &item{
				val:      tt.fields.val,
				deadline: tt.fields.deadline,
			}
			if got := itm.deadLineBefore(tt.args.t); got != tt.want {
				t.Errorf("deadLineBefore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalCacheLoop(t *testing.T) {
	cnt := 0
	localCache := NewBuildLocalCache(100, time.Second, func(cache *LocalCache) {
		cnt++
	})
	err := localCache.Set(context.Background(), "key1", "key1", time.Second)
	require.NoError(t, err)
	time.Sleep(3 * time.Second)
	localCache.RLock()
	defer localCache.RUnlock()
	_, ok := localCache.cache["key1"]
	require.False(t, ok)
	require.Equal(t, 1, cnt)
}
