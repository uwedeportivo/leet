package leet

import "testing"

func TestIsInterleave(t *testing.T) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"

	expected := true
	res := IsInterleave(s1, s2, s3)

	if expected != res {
		t.Errorf("expected %t, got %t", expected, res)
	}

	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbbaccc"

	expected = false
	res = IsInterleave(s1, s2, s3)

	if expected != res {
		t.Errorf("expected %t, got %t", expected, res)
	}
}
