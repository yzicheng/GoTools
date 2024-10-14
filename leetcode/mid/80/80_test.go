package _0

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
	}{
		{
			name:  "case 1",
			nums1: []int{1, 1, 1},
		},
		{
			name:  "case 2",
			nums1: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceLen := removeDuplicates(tt.nums1)
			fmt.Println(tt.nums1)
			fmt.Println(sliceLen)
		})
	}
}
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	fmt.Println(nums)
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			fmt.Println(nums)
			slow++
		}
		fast++
	}
	return slow
}
