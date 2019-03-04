package leet

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	s := "(()"
	expected := 2

	res := LongestValidParen(s)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	s = ")()())"
	expected = 4

	res = LongestValidParen(s)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
