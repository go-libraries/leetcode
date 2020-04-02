package _03longest_substring_without_repeating_characters


func LengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	bs := []byte(s)
	l := len(bs)
	t := make([]byte, 0)
	max := 1
	for i := 0; i < l; i++ {
		lt := len(t)
		if lt == 0 {
			t = append(t, bs[i])
		} else {
			for j := 0; j < lt; j++ {
				if bs[i] == t[j] {
					t = t[j+1:]
					break
				}
			}

			t = append(t, bs[i])
		}
		lt = len(t)
		if max < lt {
			max = lt
		}
	}

	return max
}


func LengthOfLongestSubstring1(s string) int {
	val := []byte(s)
	kvMap := make([]int, 128)
	lens := len(s)
	var max, num int
	for i, j := 0, 0; i < lens && j < lens; j++ {
		if kvMap[val[j]] > i {
			i = kvMap[val[j]]
		}
		num = j - i + 1
		if num > max {
			max = num
		}
		kvMap[val[j]] = j + 1
	}
	return max
}
