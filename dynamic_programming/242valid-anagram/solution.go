package _42valid_anagram

//本来打算用排序，可是字符串还要转slice，步骤比较多
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
		counter[t[i]]--
	}

	for _, v := range counter {
		if v != 0 {
			return false
		}
	}
	return true

}
func IsAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		counter[t[i]]--
		if counter[t[i]] < 0 {
			return false
		}
	}

	return true

}
