package _43integer_break

func integerBreak(n int) int {

	//		// i = 3 j=1 dp[i] = max(0, max(j*2, dp[3-1])) => 2
	//		// i = 3 j=2 dp[i] = max(2, max(1*2, dp[3-2])) => 2
	//		// i = 4 j=1 dp[i] = max(0, max(1*3, dp[4-1])) => 3
	//		// i = 4 j=2 dp[i] = max(3, max(2*2, 1)) => 4
	//		// i = 4 j=3 dp[i] = max(4, max(1*3, 3)) => 4
	//		dp[i] = getMax(dp[i], j*getMax(i-j, dp[i-j]))

	//return dp[n]
	dp := make([]int, n + 1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = getMax(curMax, getMax(j * (i - j), j * dp[i - j]))
		}
		dp[i] = curMax
	}
	return dp[n]


}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

