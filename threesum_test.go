package leet

import "testing"

func TestThreeSum(t *testing.T) {
	expected := [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := ThreeSum(nums)

	if len(expected) != len(res) {
		t.Errorf("expected %v, got %v", expected, res)
	}

	for i, xs := range res {
		if !IntsEqual(xs, expected[i]) {
			t.Errorf("expected %v, got %v", expected[i], xs)
		}
	}
}
