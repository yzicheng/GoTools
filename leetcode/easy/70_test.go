package easy

import "testing"

func Test70(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "case 1",
			n:    2,
			want: 2,
		},
		{
			name: "case 2",
			n:    3,
			want: 3,
		},
		{
			name: "case 3",
			n:    4,
			want: 5,
		},
		{
			name: "case 4",
			n:    5,
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := climbStairs(tt.n)
			if l != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", l, tt.want)
			}
		})
	}
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-2) + climbStairs(n-1)
}
