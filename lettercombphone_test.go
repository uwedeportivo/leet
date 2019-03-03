package leet

import (
	"sort"
	"testing"
)

func StringsEqual(as, bs []string) bool {
	if len(as) != len(bs) {
		return false
	}

	if len(as) == 0 {
		return true
	}
	for i, v := range as {
		if bs[i] != v {
			return false
		}
	}
	return true
}

func StringsEqualNoOrder(as, bs []string) bool {
	if len(as) != len(bs) {
		return false
	}

	if len(as) == 0 {
		return true
	}

	sort.Strings(as)
	sort.Strings(bs)
	for i, v := range as {
		if bs[i] != v {
			return false
		}
	}
	return true
}

func TestLetterCombinations(t *testing.T) {
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	digits := "23"

	res := LetterCombinations(digits)

	if !StringsEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
