package leet

import (
	"math"
)

func atConcat(nums1, nums2 []int, ix int) float64 {
	if ix < len(nums1) {
		return float64(nums1[ix])
	} else {
		return float64(nums2[ix-len(nums1)])
	}
}

func median(xs []int) float64 {
	n := len(xs)
	return (float64(xs[(n-1)/2]) + float64(xs[n/2])) / 2.0
}

// medianConcat returns the median of nums1 and nums2 if they do not overlap
// in their value ranges and nums1 values are smaller than nums2 values.
// If those conditions don't hold, returns false as the second return value.
func medianConcat(nums1, nums2 []int) (float64, bool) {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0, true
	}

	if len(nums1) == 0 {
		return median(nums2), true
	}

	if len(nums2) == 0 {
		return median(nums1), true
	}

	n1, n2 := len(nums1), len(nums2)
	nt := n1 + n2

	if nums1[n1-1] <= nums2[0] {
		mi := (nt - 1)/2
		if nt % 2 == 0 {
			v1 := atConcat(nums1, nums2, mi)
			v2 := atConcat(nums1, nums2, mi + 1)
			return (v1 + v2) / 2.0, true
		} else {
			return atConcat(nums1, nums2, mi), true
		}
	}
	return 0.0, false
}

func atLeftOfCut(xs []int, ix int) int {
	if ix == 0 {
		return math.MinInt32
	} else {
		return xs[(ix - 1) / 2]
	}
}

func atRightOfCut(xs []int, ix int) int {
	if ix == len(xs) * 2 {
		return math.MaxInt32
	} else {
		return xs[ix / 2]
	}
}

/*
  https://leetcode.com/problems/median-of-two-sorted-arrays/
       discuss/2471/very-concise-ologminmn-iterative-solution-with-detailed-explanation:
*/
func findMedianSortedArraysImpl(nums1, nums2 []int) float64 {
	if len(nums2) > len(nums1) {
		return findMedianSortedArraysImpl(nums2, nums1)
	}

	n1, n2 := len(nums1), len(nums2)
	nt := n1 + n2

	var l1, l2, r1, r2 int

	lo, hi := 0, n2 * 2

	for lo <= hi {
		mid2 := (lo + hi) / 2
		mid1 := nt - mid2

		l1, l2 = atLeftOfCut(nums1, mid1), atLeftOfCut(nums2, mid2)
		r1, r2 = atRightOfCut(nums1, mid1), atRightOfCut(nums2, mid2)

		if l1 > r2 {
			lo = mid2 + 1
		} else if l2 > r1 {
			hi = mid2 - 1
		} else {
			return (math.Max(float64(l1), float64(l2)) + math.Min(float64(r1), float64(r2))) / 2.0
		}
	}

	return -1.0

}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	v, ok := medianConcat(nums1, nums2)
	if ok {
		return v
	}

	v, ok = medianConcat(nums2, nums1)
	if ok {
		return v
	}

	return findMedianSortedArraysImpl(nums1, nums2)
}
