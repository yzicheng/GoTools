package _5

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
			nums1: []int{2, 3, 1, 1, 4},
		},
		{
			name:  "case 1",
			nums1: []int{3, 2, 1, 0, 4},
		},
		{
			name:  "case 1",
			nums1: []int{7, 6, 4, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceLen := canJump(tt.nums1)
			fmt.Println(tt.nums1)
			fmt.Println(sliceLen)
		})
	}
}

//给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
//
//
//
//示例 1：
//
//输入：nums = [2,3,1,1,4]
//输出：true
//解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
//示例 2：
//
//输入：nums = [3,2,1,0,4]
//输出：false
//解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。

func canJump(nums []int) bool {
	// 能够跳到最终位置，每个下标需要的距离
	index := len(nums) - 1
	for index > 0 {
		for i := 0; i < index; i++ {
			if nums[i] >= index-i {
				index = i
				break
			} else {
				if i == index-1 {
					return false
				}
			}
		}
	}
	return true
}

//给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
//
//每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
//
//0 <= j <= nums[i]
//i + j < n
//返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
//
//
//
//示例 1:
//
//输入: nums = [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//示例 2:
//
//输入: nums = [2,3,0,1,4]
//输出: 2

func jump(nums []int) int {
	// 能够跳到最终位置，每个下标需要的距离
	index := len(nums) - 1
	step := 0
	for index > 0 {
		for i := 0; i < index; i++ {
			if nums[i] >= index-i {
				step++
				index = i
				break
			}
		}
	}
	return step
}
