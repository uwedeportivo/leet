package leet

import "testing"

func TestMaximalRectangle(t *testing.T) {
	matrix := [][]byte{
		{1,0,1,0,0},
		{1,0,1,1,1},
		{1,1,1,1,1},
		{1,0,0,1,0},
	}

	expected := 6

	res := MaximalRectangle(matrix)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	matrix = [][]byte{
		{1},
	}

	expected = 1

	res = MaximalRectangle(matrix)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	matrix = [][]byte{
		{1, 1},
	}

	expected = 2

	res = MaximalRectangle(matrix)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
