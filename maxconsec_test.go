package leet

import "testing"

func TestLongestOnes(t *testing.T) {
	nums := []int{1,1,1,0,0,0,1,1,1,1,0}
	flips := 2
	expected := 6

	res := LongestOnes(nums, flips)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = []int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}
	flips = 3
	expected = 10

	res = LongestOnes(nums, flips)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
