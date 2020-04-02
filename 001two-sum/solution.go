package _01two_sum

func TwoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+nums[i] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}

}

func TwoSum2(nums []int, target int) []int {
	maps := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		maps[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		index := target - nums[i]
		if val, ok := maps[index]; ok {
			return []int{i, val}
		}
	}
	return []int{}
}

func TwoSum3(nums []int, target int) []int {
	maps := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		index := target - nums[i]
		if val, ok := maps[index]; ok {
			return []int{i, val}
		}
		maps[nums[i]] = i
	}
	return []int{}
}

//最粗暴的方式 时间复杂度 O(logN)
func TwoSum5(nums []int, target int) []int {

	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

//额外空间 复杂度 O(n) 时间复杂度 O(n)
func TwoSum4(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}