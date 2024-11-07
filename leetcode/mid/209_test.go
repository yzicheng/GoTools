package mid

import (
	"testing"
)

// 209. 长度最小的子数组
// 中等
// 相关标签
// 相关企业
// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其总和大于等于 target 的长度最小的
// 子数组
// [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//
// 示例 1：
//
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 示例 2：
//
// 输入：target = 4, nums = [1,4,4]
// 输出：1
// 示例 3：
//
// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0
func Test209(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "case 1",
			nums:   []int{2, 3, 1, 2, 4, 3},
			target: 7,
			want:   2,
		},
		{
			name:   "case 2",
			nums:   []int{1, 4, 4},
			target: 4,
			want:   1,
		},
		{
			name:   "case 3",
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			target: 11,
			want:   0,
		},
		{
			name:   "case 4",
			nums:   []int{1, 2, 3, 4, 5},
			target: 11,
			want:   3,
		},
		{
			name:   "case 5",
			nums:   []int{1, 1, 1, 1, 7},
			target: 7,
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := minSubArrayLen(tt.target, tt.nums)
			if l != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", l, tt.want)
			}
		})
	}
}

// 暴力超时
func minSubArrayLenv1(target int, nums []int) int {
	res := 0
	left := 0
	for left < len(nums) {
		sum := nums[left]
		if sum >= target {
			if res == 0 {
				res = 1
			} else {
				res = min(res, 1)
			}
			left++
			continue
		}
		j := left + 1
		for j < len(nums) {
			sum += nums[j]
			if sum >= target {
				if res == 0 {
					res = j - left + 1
				} else {
					res = min(res, j-left+1)
				}
				break
			} else {
				j++
			}
		}
		left++
	}
	return res
}

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	res := 0
	sum := nums[left]
	for left < len(nums) {
		for right < len(nums) {
			if sum >= target {
				if res == 0 {
					res = right - left + 1
				} else {
					res = min(res, right-left+1)
				}
				break
			}
			right++
			if right < len(nums) {
				sum += nums[right]
			}
		}
		sum -= nums[left]
		left++

	}
	return res
}
