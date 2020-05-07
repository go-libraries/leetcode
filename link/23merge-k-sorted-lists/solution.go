package _3merge_k_sorted_lists

import "container/heap"

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


func mergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	// 连上l1剩余的链,连上l2剩余的链
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dummy.Next
}

// Time: O(k*n), Space: O(1) n是所有链表的节点总数，k是链表个数
func mergeKLists1(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil // 链表为空或长度为0，直接返回空指针
	}
	var result *ListNode // 否则定义合并后的结果链表
	for _, list := range lists { // 定义链表数组
		// 和数组中的链表一个个合并
		result = mergeTwoSortedLists(result, list)
	}
	return result // 最后返回合并后的链表
}


func mergeKLists(lists []*ListNode) *ListNode {

	l := merge(lists)
	return l
}

func merge(lists []*ListNode) (l *ListNode) {
	if len(lists) == 0 {
		return nil
	}
	l = &ListNode{}
	min := 2147483647
	minPos := -1
	for i, v := range lists {
		if v != nil {
			if min > v.Val {
				minPos = i
				min = v.Val
			}
		}
	}

	if minPos == -1 {
		return nil
	}
	//l.Next = &ListNode{}
	l.Val = min
	if lists[minPos].Next != nil {
		lists[minPos] = lists[minPos].Next
	} else {
		if minPos != 0 {
			lists = append(lists[0:minPos], lists[minPos+1:]...)
		} else {
			lists = lists[1:]
		}
	}
	l.Next = merge(lists)
	return l
}


type PQ []*ListNode

func (p PQ) Len() int { return len(p) }
func (p PQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PQ) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}

func (p *PQ) Push(x interface{}) {
	node := x.(*ListNode)
	*p = append(*p, node)
}

func (p *PQ) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}

func mergeKLists2(lists []*ListNode) *ListNode {
	h := &ListNode {
		Val: -1,
		Next: nil,
	}
	t := h
	if len(lists) == 0 {
		return h.Next
	}

	pq := make(PQ, 0)
	for i, _ := range lists {
		if lists[i] != nil {
			pq = append(pq, lists[i])
		}
	}
	heap.Init(&pq)

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*ListNode)
		next := item.Next

		item.Next = t.Next
		t.Next = item
		t = item

		if next != nil {
			heap.Push(&pq, next)
		}
	}

	return h.Next

}

