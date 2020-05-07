package _83minimum_cost_for_tickets
/**
我们用 dp(i) 来表示从第 ii 天开始到一年的结束，我们需要花的钱。考虑到一张通行证可以让我们在「接下来」的若干天进行旅行，所以我们「从后往前」倒着进行动态规划。

对于一年中的任意一天：

如果这一天不是必须出行的日期，那我们可以贪心地选择不买。这是因为如果今天不用出行，那么也不必购买通行证，并且通行证越晚买越好。所以有 dp(i)=dp(i+1)；

如果这一天是必须出行的日期，我们可以选择买 1，7 或 30 天的通行证。
若我们购买了 j 天的通行证，那么接下来的 j - 1j−1 天，我们都不再需要购买通行证，只需要考虑第 i+j 天及以后即可。因此，我们有

dp(i)=min{cost(j)+dp(i+j)},j∈{1,7,30}

其中 cost(j) 表示 j 天通行证的价格。为什么我们只需要考虑第i+j 天及以后呢？
这里和第一条的贪心思路是一样的，如果我们需要购买通行证，那么一定越晚买越好，在握着一张有效的通行证的时候购买其它的通行证显然是不划算的。

由于我们是倒着进行动态规划的，因此我们可以使用记忆化搜索，减少代码的编写难度。
我们使用一个长度为 366的数组（因为天数是 [1,365]，而数组的下标是从 0 开始的）存储所有的动态规划结果，
这样所有的 dp(i) 只会被计算一次（和普通的动态规划相同），时间复杂度不会增大。

最终的答案记为 dp(1)。

 */
func mincostTickets1(days []int, costs []int) int {
	memo := [366]int{}  //365天
	dayM := map[int]bool{} //出行的天
	for _, d := range days {
		dayM[d] = true
	}

	var dp func(day int) int
	dp = func(day int) int {
		if day > 365 {
			return 0
		}
		if memo[day] > 0 {
			return memo[day]
		}
		if dayM[day] {
			memo[day] = min1(min1(dp(day + 1) + costs[0], dp(day + 7) + costs[1]), dp(day + 30) + costs[2])
		} else {
			memo[day] = dp(day + 1)
		}
		return memo[day]
	}
	// dp(i) = min(min1(dp(day + 1) + costs[0], dp(day + 7) + costs[1]),  dp(day + 30) + costs[2])
	// 第i天 的最低出行价格为 （（第i+1天的价格 和 第i+7天价格对比） 的最小值 对比 第i+30天价格） 的最小值
	// dp(i)=min{cost(j)+dp(i+j)},j∈{1,7,30}
	// dp(1) = min(cost(0)+dp(1+1), cost(1)+dp(1+7), cost(2)+dp(1+30))
	// => dp(335) = min(cost(0)+dp(336), cost(1)+dp(342), cost(2)+dp(365))
	return dp(1)
}

func min1(x, y int) int {
	if x < y {
		return x
	}
	return y
}

