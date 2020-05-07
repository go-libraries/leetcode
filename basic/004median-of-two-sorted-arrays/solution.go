package _04median_of_two_sorted_arrays


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l1 := len(nums1)
	l2 := len(nums2)
	la := l1+l2
	all := make([]int, la)

	if l1 > 0 {
		all[0] = nums1[0]
	}else{
		all[0] = nums2[0]
	}

	j:=0
	k:=0
	for i:=0; i < la; i ++ {
		if j >= l1 {
			if k < l2 {
				all[i] = nums2[k]
				k++
				continue
			}else{
				break
			}
		}

		if k >= l2 {
			if j < l1 {
				all[i] = nums1[j]
				j++
				continue
			}else{
				break
			}
		}

		if nums1[j] > nums2[k] {
			all[i] = nums2[k]
			k++
		}else{
			all[i] = nums1[j]
			j++
		}


	}
	middle := la / 2
	if la % 2 == 1 {
		return float64(all[middle])
	}
	return float64(all[middle] + all[middle-1]) / 2
}
