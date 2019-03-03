package leet

import "testing"

func IntsEqual(as, bs []int) bool {
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

func TestTwoSums(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}

	res := TwoSum(nums, target)

	if !IntsEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
