package _60subarray_sum_equals_k

import "testing"

func TestSubArraySum(t *testing.T) {
	//a := subarraySum([]int{1,1,1}, 2)
	//t.Log(a)
	//if a != 2 {
	//	t.Fatalf("subarray fail")
	//}

	b := subarraySum([]int{2,3,1,1,4}, 5)
	t.Log(b)

	if b != 3 {
		t.Fatalf("subarray fail")
	}

	c := subarraySum([]int{-1,-1, 1}, 0)
	t.Log(c)

	if c != 1 {
		t.Fatalf("subarray fail")
	}
}
