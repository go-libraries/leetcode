package _26remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
