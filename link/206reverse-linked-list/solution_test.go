package _06reverse_linked_list

import (
	"testing"
)

func TestReverseList(t *testing.T) {

	var node1, node2, node3, node4 *ListNode
	node1 = &ListNode{1, node2}
	node2 = &ListNode{2, node3}
	node3 = &ListNode{3, node4}
	node4 = &ListNode{4, nil}
	t.Log(ReverseList(node1))
	t.Log(ReverseList2(node1))
}
