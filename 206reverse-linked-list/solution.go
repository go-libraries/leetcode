package _06reverse_linked_list

/**
* Definition for singly-linked list.
*
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev, cur *ListNode
	cur = head
	prev = nil
	for {
		cur.Next, prev, cur = prev, cur, cur.Next
		if cur == nil {
			break
		}
	}
	return prev
}
func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := ReverseList2(head.Next)
	head.Next.Next, head.Next = head, nil
	return p

}
