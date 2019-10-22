package _42linked_list_cycle_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
	ptr1, ptr2 := head, slow.Next
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return ptr1
}

func detectCycle2(head *ListNode) *ListNode {
	m := make(map[*ListNode]int)
	cur := head
	for cur != nil {
		if _, exist := m[cur]; exist {
			return cur
		}
		m[cur] = 1
		cur = cur.Next
	}
	return nil
}
