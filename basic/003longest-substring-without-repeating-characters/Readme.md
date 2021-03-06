## 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

>输入: "abcabcbb"
>
>输出: 3
> 
>解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:

> 输入: "bbbbb"
>
> 输出: 1
>
> 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:

> 输入: "pwwkew"
>
> 输出: 3
>
> 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

## 题解：

如果是空字符串 返回0 （边界值问题）
如果是要鉴定中文 =》  rune

### 解题思路1：

    如果是空字符串 返回0 （边界值问题）
    临时存储数组 t  完整数组 bs 如果是要鉴定中文 =》  rune
    对完整数组循环:  t作为子串 有则循环比较 无则append 如果比较出结果为一致的情况 则切割数组 重新分配临时数组
    每次循环 都和上一次max比较 当大于的时候 max 等于当前长度

### 解题思路2：

    用一个map存储出现的字符的位置
    如果遇到重复的字符，则计算中间距离（即 下标差距）
    每次计算 都和最大距离比较  当大于时替换