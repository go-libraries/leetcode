package _45AddTwoNumbersII

type ListNode struct {
	Val  int
	Next *ListNode
}

//1. 读取所有值到数组中
//2. 从数组最高位开始相加，超过十则记标记为+1
//3. 重写链表
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var i1 []byte
	var i2 []byte
	//var b1end bool
	//var b2end bool
	////遍历写入数组
	//for {
	//	if !b1end {
	//		i1 = append(i1, byte(l1.Val))
	//		if l1.Next == nil {
	//			b1end = true
	//		}
	//		l1 = l1.Next
	//	}
	//	if !b2end {
	//		i2 = append(i2, byte(l2.Val))
	//		if l2.Next == nil {
	//			b2end = true
	//		}
	//		l2 = l2.Next
	//	}
	//	if b1end && b2end {
	//		break
	//	}
	//}
	for {
		i1 = append(i1, byte(l1.Val))
		if l1.Next == nil {
			break
		}
		l1 = l1.Next
	}

	for {
		i2 = append(i2, byte(l2.Val))
		if l2.Next == nil {
			break
		}
		l2 = l2.Next
	}
	var i = 1
	var up = false
	ll1 := len(i1)
	ll2 := len(i2)
	var l3 *ListNode
	for {
		if ll1 - i < 0 && ll2 -i < 0 {
			break
		}
		l4 := ListNode{
			Val: 0,
			Next: l3,
		}
		if up {
			l4.Val = l4.Val + 1
			up = false
		}
		if ll1-i > -1 {
			l4.Val = l4.Val + int(i1[ll1-i])
		}
		if ll2-i > -1 {
			l4.Val = l4.Val + int(i2[ll2-i])
		}
		if l4.Val > 9 {
			l4.Val -= 10
			up = true
		}
		l3 = &l4
		i++
	}
	if up {
		l3 = &ListNode{
			Val:  1,
			Next: l3,
		}
	}
	return l3
}


func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var i1 []byte
	var i2 []byte
	for {
		i1 = append(i1, byte(l1.Val))
		if l1.Next == nil {
			break
		}
		l1 = l1.Next
	}

	for {
		i2 = append(i2, byte(l2.Val))
		if l2.Next == nil {
			break
		}
		l2 = l2.Next
	}

	l := len(i1)
	ll2 := len(i2)
	bt := i1
	bt2 := i2
	if l < ll2 {
		l = ll2
		bt = i2
		bt2 = i1
		ll2 = len(i1)
	}
	//fmt.Println(l,  int(s1[0] - 48))
	l3 := &ListNode{
		Val: int(bt[l-1]),
	}
	var up = false

	l3.Val = int(bt2[ll2-1]) + l3.Val
	if l3.Val > 9 {
		l3.Val -= 10
		up = true
	}

	for i := 2; l-i > -1; i++ {
		l4 := ListNode{
			Val:  int(bt[l-i]),
			Next: l3,
		}
		if up {
			l4.Val = l4.Val + 1
			up = false
		}
		//fmt.Println(bt, up, l3.Val)
		// 7 2 4 3
		//   5 6 4
		// 7 8 0 7
		//fmt.Println(l4.Val, ll2-i, l-i, i,l)
		if ll2-i > -1 {
			l4.Val = l4.Val + int(bt2[ll2-i])
		}
		if l4.Val > 9 {
			l4.Val -= 10
			up = true
		}

		l3 = &l4
	}
	if up {
		l3 = &ListNode{
			Val:  1,
			Next: l3,
		}
	}
	return l3
}