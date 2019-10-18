package _24swap_nodes_in_pairs

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := &ListNode{0, head}
	hint := prev
	for prev.Next != nil && prev.Next.Next != nil {
		prev.Next.Next.Next, prev.Next.Next, prev.Next, prev = prev.Next, prev.Next.Next.Next, prev.Next.Next, prev.Next
	}
	return hint.Next
}
