package _248count_number_of_nice_subarrays

func getSmallestString(n int, k int) string {
	// 96 + i
	// z 122
	if k < n {
		return ""
	}
	res := make([]byte, n)
	for i, _ := range res {
		res[i] = 'a'
	}
	k = k - n
	for n > 0 {
		if k == 0 {
			break
		}
		n = n - 1
		if k > 25 {
			res[n] = uint8(122)
			k = k - 25
		} else {
			res[n] = 'a' + byte(k)
			k = 0
		}
	}
	//for i, v := range res {
	//	fmt.Println(v)
	//	if v > 0 {
	//		break
	//	}
	//	res[i] = 1
	//	fmt.Println(res[i])
	//}
	return string(res)
}
