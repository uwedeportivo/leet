package leet

func LongestOnes(nums []int, flips int) int {
	suffixLens := make([]int, flips + 1)
	zeros := make([]int, flips + 1)

	max := 0

	for _, v := range nums {
		if v == 1 {
			for i, sl := range suffixLens {
				suffixLens[i] = sl + 1
				if max < sl + 1 {
					max = sl + 1
				}
			}
		} else {
			for i, sl := range suffixLens {
				zeros[i] = zeros[i] + 1

				if zeros[i] > i {
					if i > 0 {
						suffixLens[i] = 1
						zeros[i] = 1
						if max < 1 {
							max = 1
						}
					} else {
						suffixLens[i] = 0
						zeros[i] = 0
					}
				} else {
					suffixLens[i] = sl + 1
					if max < sl + 1 {
						max = sl + 1
					}
				}
			}
		}
	}
	return max
}
