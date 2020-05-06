package _83minimum_cost_for_tickets

import "math"
/**
思路

方法一需要遍历一年中所有的天数，无论 \textit{days}days 的长度是多少。

但是观察方法一的递推式，我们可以看到，如果我们查询 \textit{dp}(i)dp(i)，而第 ii 天我们又不需要出行的话，那么 \textit{dp}dp 函数会一直向后计算 \textit{dp}(i + 1) = \textit{dp}(i + 2) = \textit{dp}(i + 3)dp(i+1)=dp(i+2)=dp(i+3) 一直到一年结束或者有一天我们需要出行为止。那么我们其实可以直接跳过这些不需要出行的日期，直接找到下一个需要出行的日期。

算法

现在，我们令 \textit{dp}(i)dp(i) 表示能够完成从第 \textit{days}[i]days[i] 天到最后的旅行计划的最小花费（注意，不再是第 ii 天到最后的最小花费）。令 j_1j
1
​
  是满足 \textit{days}[j_1] >= \textit{days}[i] + 1days[j
1
​
 ]>=days[i]+1 的最小下标，j_7j
7
​
  是满足 \textit{days}[j_7] >= \textit{days}[i] + 7days[j
7
​
 ]>=days[i]+7 的最小下标， j_{30}j
30
​
  是满足 \textit{days}[j_{30}] >= \textit{days}[i] + 30days[j30
 ]>=days[i]+30 的最小下标，那么就有：
dp(i)=min(dp(j1)+costs[0],dp(j7)+costs[1],dp(j30)+costs[2])
 */
func mincostTickets2(days []int, costs []int) int {
	memo := [366]int{}
	durations := []int{1, 7, 30}

	var dp func(idx int) int
	dp = func(idx int) int {
		if idx >= len(days) {
			return 0
		}
		if memo[idx] > 0 {
			return memo[idx]
		}
		memo[idx] = math.MaxInt32
		j := idx
		for i := 0; i < 3; i++ {
			for ; j < len(days) && days[j] < days[idx] + durations[i]; j++ { }
			memo[idx] = min(memo[idx], dp(j) + costs[i])
		}
		return memo[idx]
	}
	return dp(0)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
