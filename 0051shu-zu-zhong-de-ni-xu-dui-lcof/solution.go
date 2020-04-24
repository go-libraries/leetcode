package _051shu_zu_zhong_de_ni_xu_dui_lcof

func reversePairs(nums []int) int {
	var v = 0

	l := len(nums)
	for i:=0;i<l-1;i++ {
		for j:=i+1;j<l;j++ {
			if nums[i] > nums[j] {
				v ++
			}
		}
	}

	return v
}

func reversePairs1(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end - start)/2
	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid + 1, end)
	tmp := []int{}
	i, j := start, mid + 1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			cnt += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += end - (mid + 1) + 1
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i - start]
	}
	return cnt
}

