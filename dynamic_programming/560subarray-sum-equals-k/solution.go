package _60subarray_sum_equals_k

/**
我们可以基于方法一利用数据结构进行进一步的优化，我们知道方法一的瓶颈在于对每个 ii，我们需要枚举所有的 jj 来判断是否符合条件，这一步是否可以优化呢？答案是可以的。

我们定义 pre[i] 为 [0..i][0..i] 里所有数的和，则 pre[i] 可以由pre[i−1] 递推而来，即：

pre[i]=pre[i−1]+nums[i]

那么「[j..i] 这个子数组和为 k 」这个条件我们可以转化为

pre[i]−pre[j−1]==k

简单移项可得符合条件的下标 jj 需要满足

pre[j−1]==pre[i]−k

所以我们考虑以 ii 结尾的和为 kk 的连续子数组个数时只要统计有多少个前缀和为 -kpre[i]−k 的pre[j] 即可。我们建立哈希表 mp，以和为键，出现次数为对应的值，记录 pre[i] 出现的次数，
从左往右边更新 mp 边计算答案，那么以 ii 结尾的答案 mp[pre[i]−k] 即可在 O(1)O(1) 时间内得到。最后的答案即为所有下标结尾的和为 kk 的子数组个数之和。

需要注意的是，从左往右边更新边计算的时候已经保证了mp[pre[i]−k] 里记录的 pre[j] 的下标范围是 i0≤j≤i 。同时，由于pre[i] 的计算只与前一项的答案有关，
因此我们可以不用建立 pre 数组，直接用 pre 变量来记录 pre[i-1] 的答案即可。

 */
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre - k]; ok {
			count += m[pre - k]
		}
		m[pre] += 1
	}
	return count
}


/**
考虑以 ii 结尾和为 kk 的连续子数组个数，我们需要统计符合条件的下标 jj 的个数，其中  i0≤j≤i 且 [j..i][j..i] 这个子数组的和恰好为 k。

我们可以枚举 [0..i][0..i] 里所有的下标 j 来判断是否符合条件，可能有读者会认为假定我们确定了子数组的开头和结尾，还需要 O(n) 的时间复杂度遍历子数组来求和，那样复杂度就将达到 O(n^3)
3
 ) 从而无法通过所有测试用例。但是如果我们知道 [j,i] 子数组的和，就能 O(1) 推出 [j-1,i] 的和，因此这部分的遍历求和是不需要的，我们在枚举下标 jj 的时候已经能 O(1) 求出 [j,i] 的子数组之和。
 */
func subarraySum1(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

