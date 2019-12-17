package _32implement_queue_using_stacks

import "container/list"

type MyQueue struct {
    queue *list.List
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{
    	queue:list.New(),
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.queue.PushBack(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	var ret int
	if this.Empty() == false {
		ret = this.queue.Front().Value.(int)
		this.queue.Remove(this.queue.Front())
	}
	return ret
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.queue.Front().Value.(int)
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return this.queue.Len() == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */