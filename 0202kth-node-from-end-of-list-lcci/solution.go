package _202kth_node_from_end_of_list_lcci

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

func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return -1
	}

	p1 := head
	p2 := head
	i := 0
	for {

		if p1 == nil {
			break
		}

		p1 = p1.Next
		if i >= k {
			p2 = p2.Next
		}

		i++
	}

	if i < k {
		return -1
	}
	return p2.Val
}

const IntMax = 2147483647
const IntMin = -2147483648

func reverse(x int) int {
	//x1 := int32(x)
	if x < 10 && x > -1 {
		return x
	}

	isLtZero := false
	if x < 0 {
		x = x * -1
		isLtZero = true
	}
	var i int
	var v int
	var result int
	for {
		t := i * 10
		if x < t {
			break
		}
		if t == 0 {
			v = x % 10
		} else {
			v = x / t % 10
		}
		if i == 0 {
			i = 1
			result = v
		} else {
			i = i * 10
			result = result*10 + v
		}
	}
	if isLtZero {
		result *= -1
	}
	if result > IntMax {
		return 0
	}
	if result < IntMin {
		return 0
	}

	return result
}

//25 10 5 1
//25 = 10 + 10 + 5
//10 = 5 + 5
//5 = 5
//5 = 1+1+1+1+1
func waysToChange(n int) int {
	const v int = 1000000007
	coins := []int{1, 5, 10, 25}
	dp := make([]int, n+1)
	dp[0] = 1

	// 5
	//  i=0, j = 1, coins[i] = 1
	// dp[1] = dp[1] + dp[1-1]  = 1
	// j =2    1
	// j = 3   1
	// j = 4   1
	// j = 5    1
	//  i=0, j = 1, coins[i] = 5
	// dp[5] = dp[5] + dp[5-5] = 1 + 1 = 2


	// i  = 1 j = 5
	// dp[5] = dp[5] + dp[5-1]  dp[5] + dp[4]
	//背包问题
	for i := 0; i < len(coins); i++ {   //有len个物品   随着物品下标组合次数
			// 第一次初始化结果为第一个物品的组合次数（针对每个数字）
			//比如 ：  1,5,10   n:10
			// 1 0 0 0 0 0 0 0 0 0 0
			// i=0 j = 1 j < 11
			// 1 1 1 1 1 1 1 1 1 1 1
			// i = 1 j = 5
			// 1 1 1 1 1 [1+1 1+1 1+1 1+1 1+1] 1+2
			// i = 2 j = 10
			// 1 1 1 1 1 2 2 2 2 2  [3+1]
			// 如果n <15 结果集是一样的 > 15有累计计算
		for j := coins[i]; j < n+1; j++ {
			dp[j] = (dp[j] + dp[j-coins[i]]) % v
		}

	}


	return dp[n]
}
