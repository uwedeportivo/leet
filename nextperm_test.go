package leet

import "testing"

func TestNextPermutation(t *testing.T) {
	nums := []int{1,2,3}
	expected := []int{1,3,2}

	NextPermutation(nums)

	if !IntsEqual(expected, nums) {
		t.Errorf("expected %v, got %v", expected, nums)
	}

	nums = []int{1,3,2}
	expected = []int{2,1,3}

	NextPermutation(nums)

	if !IntsEqual(expected, nums) {
		t.Errorf("expected %v, got %v", expected, nums)
	}

	nums = []int{4,2,0,2,3,2,0}
	expected = []int{4,2,0,3,0,2,2}

	NextPermutation(nums)

	if !IntsEqual(expected, nums) {
		t.Errorf("expected %v, got %v", expected, nums)
	}
}