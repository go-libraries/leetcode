package _08string_to_integer_atoi

const IntMax = 2147483647
const IntMin = -2147483648

var numbers = []uint8{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}

func myAtoi(str string) int {
	bt := []byte(str)
	isLeZero := false // 未找到
	//resbt := make([]byte,0)
	var result int
	lastValType := int8(0) //未知 1空格  2 符号 + 3符号- 4数字
	for _, v := range bt {
		switch v {
		case ' ':
			if lastValType > 1 {
				goto re
			}
			lastValType = 1
		case '+':
			if lastValType > 1 {
				goto re
			}
			lastValType = 2
		case '-':
			if lastValType > 1 {
				goto re
			}
			lastValType = 3
			isLeZero = true
		default:
			var tv int = -1
			for i, v1 := range numbers {
				if v == v1 {
					tv = i
					break
				}
			}

			if tv == -1 {
				goto re
			}
			lastValType = 4
			result = result*10 + tv
			if result > IntMax {
				goto re
			}
		}

	}
re:

	if isLeZero {
		result = result * -1
	}
	if result > IntMax {
		return IntMax
	}
	if result < IntMin {
		return IntMin
	}
	return result
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	t := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	l :=  len(nums)
	if l == 0 {
		return nil
	}

	if l == 1{
		t.Val = nums[0]
	}

	return dpTree(nums)
}

func dpTree(nums[] int) *TreeNode {
	maxPos := 0
	max := -1
	l := len(nums)
	for i:=0;i<l;i++ {
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
