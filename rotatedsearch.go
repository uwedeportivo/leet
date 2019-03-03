package leet

import "sort"

func rotatedSearchImpl(nums []int, target int, offset int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return offset
		}
		return -1
	}

	n := len(nums)
	m := n / 2

	if nums[0] <= nums[m - 1] {
		if target >= nums[0] && target <= nums[m-1] {
			index := sort.SearchInts(nums[:m], target)

			if index < m  && nums[index] == target {
				return offset + index
			}

			return -1
		} else {
			return rotatedSearchImpl(nums[m:], target, offset + m)
		}
	} else { // nums[m] <= nums[n-1]
		if target >= nums[m] && target <= nums[n-1] {
			index := sort.SearchInts(nums[m:], target)

			if index < (n - m) && nums[index + m] == target {
				return offset + m + index
			}

			return -1
		} else {
			return rotatedSearchImpl(nums[:m], target, offset)
		}
	}
}

func RotatedSearch(nums []int, target int) int {
	return rotatedSearchImpl(nums, target, 0)
}
