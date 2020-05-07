package _05longest_palindromic_substring


func longestPalindrome(s string) string {
	bs := []byte(s)
	l := len(bs)
	if l < 2 {
		return s
	}

	var start, end int
	//start =  i
	//end = j
	for i := 0; i < l; i++ {
		// aba
		//=>  0   a=a  0 - 1 < 0 break  l1 = 1
		//=>  1   b = b  => a != b l1 = 1
		//=>  2   b=b  =>  a=a => l1 = 3
		//l1 := isPalindrome(bs, i, i)
		// abba
		//=>  0   a!=b  0 - 1 < 0 break  l2 = 1
		//=>  1   b=b  =>  a = a   l2 = 4
		//l2 := isPalindrome(bs, i, i  + 1)

		l := i
		r := i
		for {
			if l < 0 || r >= len(bs) || bs[l] != bs[r] {
				break
			}
			l --
			r ++
		}
		l1 := r - l - 1

		l = i
		r = i + 1
		for {
			if l < 0 || r >= len(bs) || bs[l] != bs[r] {
				break
			}
			l --
			r ++
		}
		l2 := r - l - 1
		t := l1
		if l2 > l1 {
			t = l2
		}

		if t > end - start {
			start = i - (t - 1) / 2
			end = i + t / 2
		}
	}
	return string(bs[start:end+1])
}
