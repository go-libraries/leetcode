package offer47liwuzuidazhi

/*
**
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
输入:
[

	[1,3,1],
	[1,5,1],
	[4,2,1]

]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
*/
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// f(m)(n) = grip[m][n] + max(f(m-1)(n), f(m)(n-1))
	for i := 1; i < m; i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i][j] + max(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
