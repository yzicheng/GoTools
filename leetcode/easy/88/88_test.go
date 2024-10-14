package _8

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		m     int
		n     int
	}{
		{
			nums1: []int{0},
			nums2: []int{2, 4, 6},
			m:     0,
			n:     3,
		},
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			nums2: []int{2, 5, 6},
			m:     3,
			n:     3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			fmt.Println(tt.nums1)
		})
	}
}

// 归并算法
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	j := 0
	res := make([]int, 0, m+n)
	for i < len(nums1) || j < len(nums2) {
		if i == m {
			res = append(res, nums2[j:]...)
			break
		}
		if j == n {
			res = append(res, nums1[i:]...)
			break
		}

		if nums1[i] <= nums2[j] {
			res = append(res, nums1[i])
			i++
			continue
		} else {
			res = append(res, nums2[j])
			j++
			continue
		}
	}
	copy(nums1, res)
}
