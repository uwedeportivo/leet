package leet

import "testing"

func mergeSortedArrays(xs, ys []int) []int {
	i, j := 0, 0

	var res []int

	for {
		if xs[i] <= ys[j] {
			res = append(res, xs[i])
			i = i + 1
			if i == len(xs) {
				break
			}
		} else {
			res = append(res, ys[j])
			j = j + 1
			if j == len(ys) {
				break
			}
		}
	}

	for i < len(xs) {
		res = append(res, xs[i])
		i = i + 1
	}
	for j < len(ys) {
		res = append(res, ys[j])
		j = j + 1
	}
	return res
}

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1,2}
	nums2 := []int{3,4}

	v := FindMedianSortedArrays(nums1, nums2)

	if v != 2.5 {
		t.Logf("median of %v expected %.1f and got %.1f", mergeSortedArrays(nums1, nums2), 2.5, v)
	}

	nums1 = []int{1,2,2,2,3,6}
	nums2 = []int{2,2,2,3,4,7,8}

	v = FindMedianSortedArrays(nums1, nums2)

	if v != 2.0 {
		t.Logf("median of %v expected %.1f and got %.1f", mergeSortedArrays(nums1, nums2), 2.0, v)
	}

	nums1 = []int{1,3}
	nums2 = []int{2}

	v = FindMedianSortedArrays(nums1, nums2)

	if v != 2.0 {
		t.Logf("median of %v expected %.1f and got %.1f", mergeSortedArrays(nums1, nums2), 2.0, v)
	}

	nums1 = []int{1, 3, 5, 7, 9}
	nums2 = []int{2, 4, 6, 8}

	v = FindMedianSortedArrays(nums1, nums2)

	if v != 5.0 {
		t.Logf("median of %v expected %.1f and got %.1f", mergeSortedArrays(nums1, nums2), 5.0, v)
	}
}
