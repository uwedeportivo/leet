package leet

import (
	"sort"
	"testing"
)

func TestIndices(t *testing.T) {
	s := "barfoothefoobarman"
	w := "foo"

	res := indices(s, w)

	expected := []int{3, 9}

	if !SetEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}

	s = "aaaaaa"
	w = "aaa"

	res = indices(s, w)

	expected = []int{0, 1, 2, 3}

	if !SetEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}

	s = "cacaabaacccdaadcaabdcddbababbbddcbdcddcbbaacadaadbdcacdaadbbbbad"
	w = "bbd"

	res = indices(s, w)

	expected = []int{28}

	if !SetEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func SetEqual(xs, ys []int) bool {
	sort.Ints(xs)
	sort.Ints(ys)

	if len(xs) != len(ys) {
		return false
	}

	for i := 0; i < len(xs); i++ {
		if xs[i] != ys[i] {
			return false
		}
	}
	return true
}

func TestFindSubstring(t *testing.T) {
	s := "barfoothefoobarman"

	words := []string{"foo","bar"}

	expected := []int{0, 9}

	res := FindSubstring(s, words)

	if !SetEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}

	s = "wordgoodgoodgoodbestword"
	words = []string{"word","good","best","word"}

	expected = []int{}

	res = FindSubstring(s, words)

	if !SetEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}

	s = "aaaaaaaa"
	words = []string{"aa","aa","aa"}

	expected = []int{0, 1, 2}

	res = FindSubstring(s, words)

	if !SetEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}

	s = "a"
	words = []string{"a"}

	expected = []int{0}

	res = FindSubstring(s, words)

	if !SetEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}

	s = "aaaaaa"
	words = []string{"aaa","aaa"}

	expected = []int{0}

	res = FindSubstring(s, words)

	if !SetEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}

	s = "cacaabaacccdaadcaabdcddbababbbddcbdcddcbbaacadaadbdcacdaadbbbbad"
	words = []string{"bbd","dcb","dcd"}

	expected = []int{28}

	res = FindSubstring(s, words)

	if !SetEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
