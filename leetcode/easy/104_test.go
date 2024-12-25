package easy

import (
	"math"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test104(t *testing.T) {
	tests := []struct {
		name string
		s    *TreeNode
		want int
	}{
		{
			name: "case 1",
			s:    &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := maxDepth(tt.s)
			if l != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", l, tt.want)
			}
		})
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return 1 + int(math.Max(float64(left), float64(right)))
}
