https://leetcode-cn.com/problems/reverse-linked-list
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        if not nums: return []
        window,res = [],[]
        for i,x in enumerate(nums):
            if i >= k and window[0] <= i-k:
                window.pop(0)
            while window and nums[window[-1]] <= x:
                window.pop()
            window.append(i)
            if  i>= k-1
                res.append(nums[window[0]])
        return res