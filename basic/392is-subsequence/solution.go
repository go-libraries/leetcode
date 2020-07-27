package _92is_subsequence

//tips:  子序列  则是说s中的所有字符都会按照顺序出现在t中
//因此 采用双指针
func isSubsequence(s string, t string) bool {
	sb := []byte(s)
	tb := []byte(t)

	var i = 0
	var j = 0
	var sl = len(sb)
	var tl = len(tb)
	for ; i < sl && j < tl; {
		if sb[i] == tb[j] {
			i++
		}
		j++
	}

	if i == sl {
		return true
	}

	return false
}

