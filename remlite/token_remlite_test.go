package remlite

import (
	"fmt"
	"testing"
	"time"
)

func TestTokenLimit_Wait(t *testing.T) {
	tests := []struct {
		name       string
		duration   time.Duration
		size       int
		tokenLimit *TokenLimit
	}{
		{
			name:       "0 Size",
			duration:   time.Second,
			size:       0,
			tokenLimit: NewBuildTokenLimit(time.Second, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				go func() {

					time.Sleep(time.Second * 20)
					fmt.Println(tt.tokenLimit)
				}()
			}
			time.Sleep(time.Hour)

		})
	}
}
