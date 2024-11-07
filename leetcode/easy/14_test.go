package easy

import (
	"strings"
	"testing"
)

func Test14(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		want string
	}{
		{
			name: "case 1",
			s:    []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "case 2",
			s:    []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "case 3",
			s:    []string{"", "b"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := longestCommonPrefix(tt.s)
			if l != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", l, tt.want)
			}
		})
	}
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	index := 1
	for index < len(strs) {
		i := 0
		j := 0
		str := ""
		for i < len(prefix) && j < len(strs[index]) {
			if prefix[i] == strs[index][j] {
				str += string(prefix[i])
			} else {
				str += " "
			}
			i++
			j++
		}
		if str == "" {
			return ""
		} else {
			prefix = strings.Split(str, " ")[0]
		}
		index++
	}
	return prefix
}
