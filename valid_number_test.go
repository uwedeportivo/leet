package leet

import "testing"

func TestValidNumber(t *testing.T) {
	testCases := []struct {
		str string
		isNum bool
	}{
		{"0", true},
		{" 0.1 ", true},
		{"abc", false},
		{"1 a", false},
		{"2e10", true},
		{" -90e3   ", true},
		{" 1e", false},
		{"e3", false},
		{" 6e-1", true},
		{" 99e2.5 ", false},
		{"53.5e93", true},
		{" --6 ", false},
		{"-+3", false},
		{"95a54e53", false},
		{".1", true},
		{"+.8", true},
	}

	for _, c := range testCases {
		res := IsNumber(c.str)
		if res != c.isNum {
			t.Errorf("expected %t, got %t for string %s", c.isNum, res, c.str)
		}
	}
}
