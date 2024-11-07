package easy

import (
	"strings"
	"testing"
	"unicode"
)

//125. 验证回文串
//简单
//相关标签
//相关企业
//如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
//
//字母和数字都属于字母数字字符。
//
//给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
//
//
//
//示例 1：
//
//输入: s = "A man, a plan, a canal: Panama"
//输出：true
//解释："amanaplanacanalpanama" 是回文串。
//示例 2：
//
//输入：s = "race a car"
//输出：false
//解释："raceacar" 不是回文串。
//示例 3：
//
//输入：s = " "
//输出：true
//解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
//由于空字符串正着反着读都一样，所以是回文串。

func Test125(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "case 1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "case 2",
			s:    "race a car",
			want: false,
		},
		{
			name: "case 3",
			s:    " ",
			want: true,
		}, {
			name: "case 4",
			s:    "0P",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := isPalindrome(tt.s)
			if l != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", l, tt.want)
			}
		})
	}
}

func isPalindrome(s string) bool {
	s = removeNonAlphabetic(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func removeNonAlphabetic(s string) string {
	var result strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			result.WriteRune(r)
		}
	}
	return strings.ToLower(result.String())
}
