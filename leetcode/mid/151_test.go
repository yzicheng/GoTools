package mid

import (
	"container/list"
	"strings"
	"testing"
)

func Test151(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "case 1",
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			name: "case 2",
			s:    "  hello world  ",
			want: "world hello",
		},
		{
			name: "case 3",
			s:    "a good   example",
			want: "example good a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := reverseWords(tt.s)
			if l != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", l, tt.want)
			}
		})
	}
}

func reverseWords(s string) string {
	stack := list.New()
	temps := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if temps != "" {
				stack.PushBack(temps)
				stack.PushBack(" ")
			}
			temps = ""
		} else {
			temps += string(s[i])
			if i == len(s)-1 {
				stack.PushBack(temps)
				temps = ""
			}
		}
	}
	res := ""
	for stack.Len() > 0 {
		back := stack.Back().Value.(string)
		stack.Remove(stack.Back())
		res += back
	}
	return strings.TrimSpace(res)
}
