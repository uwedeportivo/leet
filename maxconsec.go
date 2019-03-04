package leet

func LongestOnes(nums []int, flips int) int {
	zeros := 0
	start := 0
	n := len(nums)

	max := 0

	for k, v := range nums {
		if v == 0 {
			zeros = zeros + 1

			if zeros > flips {
				if max < k - start {
					max = k - start
				}
				for start < k && nums[start] == 1 {
					start = start + 1
				}
				start = start + 1
			}
		}
	}

	if start < n {
		if max < n - start {
			max = n - start
		}
	}
	return max
}
