package _52maximum_product_subarray
func maxProduct(nums []int) int {

	l := len(nums)
	if l < 1 {
		return 0
	}
	max := nums[0]
	min := nums[0]
	res := nums[0]
	for i:=1;i<l;i++ {
		mx, mn := max, min
		max = maxf(mx * nums[i], maxf(nums[i], mn * nums[i]))
		min = minf(mn * nums[i], minf(nums[i], mx * nums[i]))

		res = maxf(max, min)
	}


	return res
}
func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx * nums[i], max(nums[i], mn * nums[i]))
		minF = min(mn * nums[i], min(nums[i], mx * nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

