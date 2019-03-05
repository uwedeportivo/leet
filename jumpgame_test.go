package leet

import "testing"

func TestJumpGame(t *testing.T) {
	nums := []int{2,3,1,1,4}
	expected := 2

	res := JumpGame(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
