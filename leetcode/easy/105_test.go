package easy

import (
	"fmt"
	"testing"
)

func Test105(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	preorder := []int{3, 9, 20, 15, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(root)
}

// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(inorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}
	// 1.找到跟节点
	// 中左右
	rootV := preorder[0]
	root := &TreeNode{
		Val: rootV,
	}
	// 找到中序遍历的根节点
	// 左中右
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	left := buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	right := buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	root.Left = left
	root.Right = right
	return root
}
