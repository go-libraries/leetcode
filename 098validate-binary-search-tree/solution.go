package _98validate_binary_search_tree

import "math"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var keys []int
	LevelOrder(root, &keys)
	for i := 1; i < len(keys); i++ {
		if keys[i-1] >= keys[i] {
			return false
		}
	}
	return true
}
func LevelOrder(node *TreeNode, keys *[]int) {
	if node == nil {
		return
	}
	LevelOrder(node.Left, keys)
	*keys = append(*keys, node.Val)
	LevelOrder(node.Right, keys)
}
func IsValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBST(root, math.MinInt64, math.MaxInt64)
}
func isBST(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}
	if left >= root.Val || right <= root.Val {
		return false
	}
	return isBST(root.Left, left, root.Val) && isBST(root.Right, root.Val, right)
}
