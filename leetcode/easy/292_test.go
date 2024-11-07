package easy

import (
	"container/list"
	"testing"
)

// 代码
// 测试用例
// 测试结果
// 测试结果
// 392. 判断子序列
// 简单
// 相关标签
// 相关企业
// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
//
// 进阶：
//
// 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
//
// 致谢：
//
// 特别感谢 @pbrother 添加此问题并且创建所有测试用例。
//
// 示例 1：
//
// 输入：s = "abc", t = "ahbgdc"
// 输出：true
// 示例 2：
//
// 输入：s = "axc", t = "ahbgdc"
// 输出：false
func Test292(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "case 1",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "case 2",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
		{
			name: "case 3",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "case 4",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "case 5",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "case 6",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "case 7",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := isSubsequence(tt.s, tt.t)
			if l != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", l, tt.want)
			}
		})
	}
}
func isSubsequence(s string, t string) bool {
	queue := list.New()
	if s == "" {
		return true
	}
	if t == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		queue.PushBack(s[i])
	}
	for i := 0; i < len(t); i++ {
		if queue.Len() == 0 {
			return true
		}
		front := queue.Front().Value.(byte)
		if front == t[i] {
			queue.Remove(queue.Front())
		}

	}
	return queue.Len() == 0
}
