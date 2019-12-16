package _25implement_stack_using_queues

import "container/list"

type MyStack struct {
	stack *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{stack: list.New()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if x == 0 {
		return
	}
	this.stack.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	var ret int
	if this.Empty() == false {
		ret = this.stack.Back().Value.(int)
		this.stack.Remove(this.stack.Back())
	}
	return ret
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.stack.Back().Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.stack.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
