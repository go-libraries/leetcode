package _56shu_zu_zhong_shu_zi_chu_xian_de_ci_shu_lcof

func singleNumbersMap(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}

	mp := make(map[int]int)
	for _, n := range nums {
		if _, ok := mp[n]; ok {
			mp[n] = 2
		} else {
			mp[n] = 1
		}
	}

	var res = make([]int, 0)
	for v, c := range mp {
		if c == 1 {
			res = append(res, v)
		}
	}

	return res
}
func singleNumbers(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}

	var result int
	for _, v := range nums {
		result ^= v
	}

	flag := result & (-result)

	// // 根据二进制位上的那个“1”进行分组
	//            // 需要注意的是，分组的结果必然是相同的数在相同的组，且还有一个结果数
	//            // 因此每组的数再与res=0一路异或下去，最终会得到那个结果数A或B
	//            // 且由于异或运算具有自反性，因此只需得到其中一个数即可
	//            if ((flag & val) != 0) {
	//                res ^= val;
	//            }
	 res :=0
	for _, num := range nums {
		if (flag & num) != 0 {
			res ^= num
		}
	}

	return []int{res, res^result}
}
