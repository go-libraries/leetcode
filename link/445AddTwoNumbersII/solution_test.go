package _45AddTwoNumbersII

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	//l1 = &ListNode{
	//	Val:  5,
	//	Next: nil,
	//}
	//l2 = &ListNode{
	//	Val:  5,
	//	Next: nil,
	//}
	l3 := addTwoNumbers(l1, l2)
	var i = 0
	for {

		fmt.Print(l3.Val, ",")
		i++
		if l3.Next == nil {
			break
		}
		l3 = l3.Next
	}
	l3 = addTwoNumbers1(l1, l2)
	fmt.Println("\n简化写法反而效率低")
	i= 0
	for {

		fmt.Print(l3.Val, ",")
		i++
		if l3.Next == nil {
			break
		}
		l3 = l3.Next
	}
	fmt.Println("")
}