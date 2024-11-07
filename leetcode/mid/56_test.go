package mid

import (
	"reflect"
	"slices"
	"testing"
)

// 以数组 intervals 表示若干个区间的集合，
// 其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//
// 示例 1：
//
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 示例 2：
//
// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]
// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
func Test56(t *testing.T) {
	tests := []struct {
		name string
		nums [][]int
		want [][]int
	}{
		{
			name: "case 1",
			nums: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		}, {
			name: "case 2",
			nums: [][]int{{4, 5}, {1, 4}},
			want: [][]int{{1, 5}},
		}, {
			name: "case 3",
			nums: [][]int{{1, 4}, {0, 4}},
			want: [][]int{{0, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := merge(tt.nums)
			if reflect.DeepEqual(tt.want, l) {
				t.Errorf("lengthOfLastWord() = %v, want %v", l, tt.want)
			}
		})
	}
}

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序
	ans := make([][]int, 0, len(intervals))
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := ans[len(ans)-1]
		// 如果last右端小于intervals的左端，说明不连续
		// 否则说明存在连续
		if last[1] < intervals[i][0] {
			ans = append(ans, intervals[i])
			continue
		} else {
			// 左端谁小取谁
			if last[0] > intervals[i][0] {
				last[0] = intervals[i][0]
			}
			// 右端谁大取谁
			if last[1] < intervals[i][1] {
				last[1] = intervals[i][1]
			}
		}
	}
	return ans
}
