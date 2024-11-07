package mid

import "testing"

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 返回容器可以储存的最大水量。
//
// 说明：你不能倾斜容器。
//
// 示例 1：
//
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
// 示例 2：
//
// 输入：height = [1,1]
// 输出：1
func Test11(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "case 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "case 2",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "case 3",
			height: []int{1, 2, 1},
			want:   2,
		},
		{
			name:   "case 4",
			height: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:   25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := maxArea(tt.height)
			if l != tt.want {
				t.Errorf("convert() = %v, want %v", l, tt.want)
			}
		})
	}
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	maxVal := 0
	for i < j {
		val := (j - i) * min(height[i], height[j])
		maxVal = max(maxVal, val)
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return maxVal
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
