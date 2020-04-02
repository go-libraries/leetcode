package _03longest_substring_without_repeating_characters

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(LengthOfLongestSubstring(""))
	fmt.Println(LengthOfLongestSubstring(" "))
	fmt.Println(LengthOfLongestSubstring("bbbbbb"))
	fmt.Println(LengthOfLongestSubstring("dvdf"))
	fmt.Println(LengthOfLongestSubstring("abcabcbb"))
}


func TestLengthOfLongestSubstring1(t *testing.T) {
	fmt.Println(LengthOfLongestSubstring1(""))
	fmt.Println(LengthOfLongestSubstring1(" "))
	fmt.Println(LengthOfLongestSubstring1("bbbbbb"))
	fmt.Println(LengthOfLongestSubstring1("dvdf"))
	fmt.Println(LengthOfLongestSubstring1("abcabcbb"))
}
