package _50intersection_of_two_arrays_ii

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	if nums1 == nil || nums2 == nil {
		return res
	}
	maps := make(map[int]int)

	for _, value := range nums1 {
		if _, ok := maps[value]; !ok {
			maps[value] = 1
		} else {
			maps[value]++
		}
	}
	for _, value := range nums2 {
		if _, ok := maps[value]; ok {
			if maps[value] > 0 {
				res = append(res, value)
			}
			maps[value]--
		}
	}
	return res
}
