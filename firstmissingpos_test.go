package leet

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	nums := []int{1,2,0}
	expected := 3

	res := FirstMissingPositive(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = []int{3,4,-1,1}
	expected = 2

	res = FirstMissingPositive(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = []int{7,8,9,11,12}
	expected = 1

	res = FirstMissingPositive(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = []int{1,1}
	expected = 2

	res = FirstMissingPositive(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = []int{-5,1000}
	expected = 1

	res = FirstMissingPositive(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = []int{2}
	expected = 1

	res = FirstMissingPositive(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = []int{4,2,0,1,4}
	expected = 3

	res = FirstMissingPositive(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = []int{1,2,3,10,2147483647,9}
	expected = 4

	res = FirstMissingPositive(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
