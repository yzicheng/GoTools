package mid

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [9,3,15,20,7] [9,15,7,20,3]
// [3,9,20,null,null,15,7]
func Test106(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
	fmt.Println(root)
}
func buildTree(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || postorder == nil {
		return nil
	}
	if len(postorder) == 0 {
		return nil
	}
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{inorder[0], nil, nil}
	}
	//找到根节点
	// postorder
	// 左右根
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}

	// inorder
	// 左根右
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	// 左边长度  len(inorder[:i])
	// 右边长度  len(inorder[i+1:])
	// 右边的位置
	rigntIndex := len(postorder) - 1 - len(inorder[i+1:])
	// 左边位置
	leftIndex := rigntIndex
	// 左
	left := buildTree(inorder[:i], postorder[:leftIndex])
	// 右
	ritht := buildTree(inorder[i+1:], postorder[leftIndex:len(postorder)-1])

	root.Left = left
	root.Right = ritht
	return root
}
