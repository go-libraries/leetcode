package _01two_sum

import "testing"

func TestTwoSum1(t *testing.T) {
	s1 := []int{2, 7, 11, 15}
	var s2, s3 []int
	copy(s2, s1)
	copy(s3, s1)
	t.Log(TwoSum1(s1, 9))
	t.Log(TwoSum2(s2, 9))
	t.Log(TwoSum3(s1, 9))
}
