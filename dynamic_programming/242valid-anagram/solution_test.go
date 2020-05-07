package _42valid_anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	s := "rat"
	ts := "tar"
	t.Log(IsAnagram(s, ts))
	t.Log(IsAnagram2(s, ts))
}
