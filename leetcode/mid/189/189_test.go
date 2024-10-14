package _89

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		k     int
		want  []int
	}{
		{
			name:  "case 1",
			nums1: []int{1, 2, 3, 4, 5, 6, 7},
			k:     3,
		},
		{
			name:  "case 2",
			nums1: []int{-1, -100, 3, 99},
			k:     4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.nums1, tt.k)
		})
		fmt.Println(tt.nums1)
	}
}

// 超时，后面还是想办法用新切片并做copy吧
func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}
	for k > 0 {
		lastValue := nums[len(nums)-1]
		for i := len(nums) - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = lastValue
		k--
	}
}

func rotate2(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}
