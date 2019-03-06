package leet

import "testing"

func TestCalculateMinimumHP(t *testing.T) {
	nums := [][]int{{-2,-3,3},{-5,-10,1},{10,30,-5}}
	expected := 7

	res := CalculateMinimumHP(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = [][]int{{1,-3,3},{0,-2,0},{-3,-3,-3}}
	expected = 3

	res = CalculateMinimumHP(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
