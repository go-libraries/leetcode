package _54maximum_binary_tree

import "testing"

func TestConstructMaximumBinaryTree(t *testing.T)  {
	constructMaximumBinaryTree([]int{})
	constructMaximumBinaryTree([]int{1,2,3,4,5})
	constructMaximumBinaryTree([]int{5,4,3,2,1})
	constructMaximumBinaryTree([]int{11,33,22,6767,3213,345345})
}