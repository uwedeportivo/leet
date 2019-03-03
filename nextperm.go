package leet

import "sort"

func NextPermutation(nums []int) {
	n := len(nums)

	for j := n-1; j >=0; j-- {
		for i := n-1; i > j; i-- {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				sort.Ints(nums[j+1:])
				return
			}
		}
	}
	sort.Ints(nums)
}


