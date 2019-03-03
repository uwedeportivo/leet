package leet

import "testing"


func TestGenerateParenthesis(t *testing.T) {
	expected := []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}

	res := GenerateParenthesis(3)

	if !StringsEqualNoOrder(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}