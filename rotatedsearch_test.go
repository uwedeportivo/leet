package leet

import "testing"

func TestRotatedSearchFound(t *testing.T) {
	nums := []int{4,5,6,7,0,1,2}
	target := 0
	expected := 4

	res := RotatedSearch(nums, target)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = []int{4,5,6,7,0,1,2}
	target = 1
	expected = 5

	res = RotatedSearch(nums, target)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}

func TestRotatedSearchNotFound(t *testing.T) {
	nums := []int{4,5,6,7,0,1,2}
	target := 3
	expected := -1

	res := RotatedSearch(nums, target)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
