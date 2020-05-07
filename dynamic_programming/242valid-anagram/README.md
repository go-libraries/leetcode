https://leetcode-cn.com/problems/valid-anagram
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        dic1,dic2 = {},{}
        for item in s:
            dic1[item] = dic1.get(item,0)+1
        for item in t:
            dic2[item] = dic2.get(item,0)+1
        return dic1==dic2

    def isAnagram2(self, s: str, t: str) -> bool:
        dic1,dic2 = [0]*26, [0]*26
        for item in s:
            dic1[ord(item)-ord('a')] +=1
        for item in t:
            dic2[ord(item)-ord('a')] +=1
        return dic1==dic2