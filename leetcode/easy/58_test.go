package easy

import (
	"strings"
	"testing"
)

func Test58(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "case 1",
			s:    "Hello World",
			want: 5,
		},
		{
			name: "case 2",
			s:    "   fly me   to   the moon  ",
			want: 4,
		},
		{
			name: "case 3",
			s:    "luffy is still joyboy",
			want: 6,
		},
		{
			name: "case 4",
			s:    "abc",
			want: 3,
		},
		{
			name: "case 4",
			s:    "",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lengthOfLastWord(tt.s)
			if l != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", l, tt.want)
			}
		})
	}
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = strings.TrimSpace(s)
	index := 1
	for index <= len(s) {
		if s[len(s)-index] != ' ' {
			index++
			continue
		} else {
			return index - 1
		}
	}
	return index - 1
}
