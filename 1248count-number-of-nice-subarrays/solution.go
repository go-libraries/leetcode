package _248count_number_of_nice_subarrays

func numberOfSubarrays(nums []int, k int) int {
	l := len(nums)
	if l < 1 {
		return 0
	}

	//oddSubLen 记录每个奇数之间的距离
	var ret, oddSubLen int
	mp := make([]int, 0)
	for i := 0; i < l; i++ {
		v := nums[i]
		oddSubLen++
		if v%2 == 1 {
			mp = append(mp, oddSubLen)
			oddSubLen = 0
		}
		if len(mp) >= k {
			ret += mp[len(mp)-k]
		}
	}
	return ret
}
