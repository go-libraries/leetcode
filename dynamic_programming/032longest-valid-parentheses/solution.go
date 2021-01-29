package _32longest_valid_parentheses

//动态规划解析此题  另外可以用栈解决
func longestValidParentheses(s string) int {
	maxAns := 0  //结果值
	dp := make([]int, len(s))  //存储每个位置的最大连续数
	//分析组合   ()()  ()  (())()
	//所以当出现 )时，可以去判断前一个符号是否是(,如果是，则符合 ()。则当前位置长度为 dp[下标-2] + 2。越界问题注意下
	//如果不是，继续往前匹配。
	//怎么匹配呢？
	//dp记录了之前一位 )的连续位数，假设是连续的，则可以忽略，直接验证 i- len(连续) - 1 是否为 （ 为 (匹配上
	//
	//如果他上一个是未连续的，自然可以忽略
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			//所以当出现 )时，可以去判断前一个符号是否是(,如果是，则符合 ()。则当前位置长度为 dp[下标-2] + 2。越界问题注意下
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i - 2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				//dp[i-1]记录是之前字符串的连续长度
				//dp记录了之前一位 )的连续位数，假设是连续的，则可以忽略，直接验证 i- len(连续) - 1 是否为 （ 为 (匹配上
				lastMathCount := dp[i - 1]
				lastMathIndex :=  i - lastMathCount
				if lastMathIndex > 0 && s[lastMathIndex-1] == '(' {  // (((()()((()))
					if lastMathIndex >= 2 { // ()(()) 0 1 2 3 4 5 i == 5   dp[i]  lastMathCount = 2 index = 3
						// 4 >= 2 => d[i] = 2 + dp[3-2] + 2 => 6  // (())(()) =>  8 = 2+ 4{dp[5-2]} + 2
						dp[i] = lastMathCount + dp[lastMathIndex - 2] + 2
					} else {
						// (())  dp[i] = 2+2 => 4
						dp[i] = lastMathCount + 2
					}
				}
			}

			maxAns = max(maxAns, dp[i])
		}
	}
	return maxAns
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
