package _20valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	//s :="()[]{"
	s1 := "{[]"
	//t.Log(IsValid(s))
	t.Log(IsValid2(s1))
}
