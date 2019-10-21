package _41linked_list_cycle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
func hasCycle2(head *ListNode) bool {
	for head != nil {
		if head == head.Next {
			return true
		}
		if head.Next != nil {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return false
}
