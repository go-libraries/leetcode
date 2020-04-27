package _3merge_k_sorted_lists

import (
	"fmt"
	"testing"
)

func TestMergeSorts(t *testing.T)  {
	//1->4->5,
	//	1->3->4,
	//	2->6
	l1 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l3 := &ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}
	l := mergeKLists1([]*ListNode{l1,l2,l3})
	for  {
		fmt.Println(l.Val)

		if l.Next == nil {
			break
		}
		l = l.Next

	}
}


