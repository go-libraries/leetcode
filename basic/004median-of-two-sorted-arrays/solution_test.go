package _04median_of_two_sorted_arrays

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	fmt.Println(findMedianSortedArrays([]int{}, []int{1,2}))
	fmt.Println(findMedianSortedArrays([]int{1,2}, []int{3,4}))
	fmt.Println(findMedianSortedArrays([]int{1,2}, []int{3}))
}


