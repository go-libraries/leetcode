package _02binary_tree_level_order_traversal

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	arr := make([][]int, 0)
	if root == nil {
		return arr
	}
	l := list.New()
	l.PushFront(root)
	for l.Len() > 0 {
		curLevel := make([]int, 0)
		listLength := l.Len()
		for i := 0; i < listLength; i++ {
			cur := l.Remove(l.Back()).(*TreeNode)
			curLevel = append(curLevel, cur.Val)
			if cur.Left != nil {
				l.PushFront(cur.Left)
			}
			if cur.Right != nil {
				l.PushFront(cur.Right)
			}
		}
		arr = append(arr, curLevel)
	}
	return arr
}
func LevelOrder2(root *TreeNode) [][]int {
	arr := [][]int{}
	if root == nil {
		return arr
	}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(arr) {
			arr = append(arr, []int{})
		}
		arr[level] = append(arr[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return arr
}
