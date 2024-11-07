package mid

import (
	"slices"
	"testing"
)

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长
// 子串
// 的长度。
//
// 示例 1:
//
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
//
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
//
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
// 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
func Test3(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "case 1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "case 2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "case 3",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "case 4",
			s:    "abba",
			want: 2,
		},
		{
			name: "case 5",
			s:    "dvdf",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lengthOfLongestSubstring(tt.s)
			if l != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", l, tt.want)
			}
		})
	}
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	left := 0
	right := 1
	sub := []byte{s[left]}
	maxVal := 1
	for left < len(s) && right < len(s) {
		if !slices.Contains(sub, s[right]) {
			sub = append(sub, s[right])
		} else {
			maxVal = max(maxVal, len(sub))
			sub = sub[slices.Index(sub, s[right])+1:]
			left = right - len(sub)
			sub = append(sub, s[right])
		}
		right++
	}
	return max(maxVal, len(sub))
}
