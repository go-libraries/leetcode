package offer48zuichangzifuchuan

// func lengthOfLongestSubstring(s string) int {

//     l := len(s)
//     if l == 0 {
//         return 0
//     }

//     res := 1
//     bts := []byte{s[0]}
//     for i:=1;i<l;i++ {
//         for j,v := range bts {
//             if v == s[i] {
//                 bts = bts[j+1:]
//                 break
//             }
//         }
//         bts = append(bts, s[i])
//         if len(bts) > res {
//                     res = len(bts)
//         }
//     }

//     return res
// }

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	res := 1
	var left, right = 0, 0 //雙指針
	mp := make(map[byte]struct{})
	for right < l {
		for left < right {
			if _, ok := mp[s[right]]; !ok { //如果右指針最後一個字符不存在，繼續遍歷右指針
				break
			}
			delete(mp, s[left]) //否則刪除，左指針右移
			left++
		}
		for right < l {
			if _, ok := mp[s[right]]; ok { //如果遭遇了同樣嘚字符  找到左邊的位置
				break
			}
			mp[s[right]] = struct{}{}
			right++
		}

		res = max(res, len(mp))
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
