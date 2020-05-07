package _83move_zeroes

func moveZeroes(nums []int) {
	var noZero []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			noZero = append(noZero, nums[i])
		}
	}
	for i := 0; i < len(noZero); i++ {
		nums[i] = noZero[i]
	}
	for i := len(noZero); i < len(nums); i++ {
		nums[i] = 0
	}
}
func moveZeroes1(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}
	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}
func moveZeroes2(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}

}
func moveZeroes3(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if k != i {
				nums[k], nums[i] = nums[i], nums[k]
			}
			k++
		}
	}

}
