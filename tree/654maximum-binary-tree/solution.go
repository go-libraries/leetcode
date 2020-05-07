package _54maximum_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}

	return dpTree(nums)
}

func dpTree(nums [] int) *TreeNode {
	maxPos := 0
	max := -1
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] > max {
			maxPos = i
			max = nums[i]
		}
	}

	t := &TreeNode{
		Val:   max,
		Left:  nil,
		Right: nil,
	}
	if maxPos != 0 {
		t.Left = dpTree(nums[0:maxPos])
	}
	if maxPos < l-1 {
		t.Right = dpTree(nums[maxPos+1:])
	}
	return t
}
