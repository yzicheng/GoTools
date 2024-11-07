package _38

import (
	"fmt"
	"reflect"
	"testing"
)

//
//给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
//题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
//
//请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
//
//
//示例 1:
//
//输入: nums = [1,2,3,4]
//输出: [24,12,8,6]
//示例 2:
//
//输入: nums = [-1,1,0,-3,3]
//输出: [0,0,9,0,0]

func Test(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		want      []int
	}{
		{
			name: "test", citations: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6},
		},
		{
			name: "test", citations: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			answer := productExceptSelf(tt.citations)
			if !reflect.DeepEqual(answer, tt.want) {
				t.Errorf("maxv !=tt.want %d,%d", answer, tt.want)
			}
			fmt.Println(answer)
		})

	}
}

func productExceptSelf(nums []int) []int {
	l := make([]int, len(nums))
	l[0] = 1
	r := make([]int, len(nums))
	r[len(r)-1] = 1
	answer := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		l[i] = nums[i-1] * l[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = nums[i+1] * r[i+1]
	}
	for i := 0; i < len(answer); i++ {
		answer[i] = l[i] * r[i]
	}
	return answer
}
