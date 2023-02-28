package dfs

/**

无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。

示例1:

 输入：S = "qwe"
 输出：["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
示例2:

 输入：S = "ab"
 输出：["ab", "ba"]
提示:

字符都是英文字母。
字符串长度在[1, 9]之间。
*/

var ans []string
var path []byte
var visit []bool

func permutation(S string) []string {
	n := len(S)
	ans = []string{}
	path = make([]byte, n)
	visit = make([]bool, n)
	var dfs func(idx int)
	dfs = func(num int) {
		if num == n {
			//如果长度等于n 添加到字符串
			ans = append(ans, string(path))
			return
		}
		for i := 0; i < n; i++ {
			if !visit[i] { //当前dfs 没有使用过 第i个字符 则加入继续迭代
				//如果当前字符没有被使用过
				//将当前字符添加到路径中继续dfs
				path[num] = S[i]
				visit[i] = true
				dfs(num + 1)
				visit[i] = false //dfs迭代完之后清除标志
				//移除掉最后一个字符串
				//path[num] = ' '
			}
		}
	}
	// num = 0   qwe      i:=0   path = q  0:true
	// num = 1    i=0  i = 0  continue  i = 1  path    q w
	dfs(0)
	return ans
}

// func dfs(num,n int, S string) {
//       if num==n{
//             //如果长度等于n 添加到字符串
//             ans = append(ans, string(path))
//             return
//         }
//         for i:=0;i<n;i++{
//             if !visit[i]{
//                 //如果当前字符没有被使用过
//                 //将当前字符添加到路径中继续dfs
//                 path[num] = S[i]
//                 visit[i]=true
//                 dfs(num+1,n, S)
//                 visit[i]=false
//                 //移除掉最后一个字符串
//                 path[num] = ' '
//             }
//         }
// }
