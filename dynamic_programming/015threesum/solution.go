package _15threesum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	l := len(nums)
	if l < 3 {
		return nil
	}
	x := make([][]int, 0)

	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			for k := i + 2; k < l; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {

					//todo : 去重即可
					x = append(x, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return x
}

func threeSum1(nums []int) [][]int {

	l := len(nums)
	if l < 3 {
		return nil
	}
	sort.Ints(nums)
	fmt.Println(nums)
	x := make([][]int, 0)
	for i := 0; i < l; i++ {

		//已排序  若当前值>0则右边的数不可能相加为0了
		if nums[i] > 0 {
			break
		}

		start := i + 1
		end := l - 1
		for {
			if start >= end {
				break
			}
			v := nums[i] + nums[start] + nums[end]
			fmt.Println(v,i,nums[i] , nums[start] , nums[end])
			if v == 0 {
				x = append(x, []int{nums[i], nums[start], nums[end]})
				for  start < end && nums[start] == nums[start+1]{
					start += 1
				}
				for start < end && nums[end] == nums[end-1]{
					end -= 1
				}
				start++
				end--
			}
			if v > 0 {
				//已排序 若当前数和上一个数一样 则结果一样 可跳过
				end = end - 1
				if end < i {
					break
				}
				if nums[end] == nums[end+1] {
					continue
				}

			} else {
				start = start + 1
				if start >= end {
					break
				}
				if nums[start] == nums[start-1] {
					continue
				}
			}
		}
	}

	return x
}
