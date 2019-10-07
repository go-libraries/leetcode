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
