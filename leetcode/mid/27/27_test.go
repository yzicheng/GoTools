package _7

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		want  int
	}{
		{
			name:  "case1",
			nums1: []int{3, 2, 2, 3},
			m:     3,
			want:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := removeElement(tt.nums1, tt.m)
			if w != tt.want {
				t.Errorf("w %v != tt.want %v", w, tt.want)
			}
			fmt.Println(tt.nums1)
			fmt.Println(w)
		})
	}
}
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
