package leet

func hole(xs []int, left, right int) int {
	if right - left == 0 {
		return left
	}

	m := (right - left) / 2 + left
	lc := 0

	for _, v := range xs {
		if v <= m && v >= left {
			lc++
		}
	}

	if m + 1 - left > lc {
		return hole(xs, left, m)
	} else {
		return hole(xs, m + 1, right)
	}
}

func FirstMissingPositive(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 1
	}

	min := 1<<7 - 1
	max := -1

	dedup := make(map[int]bool)

	xs := make([]int, 0, n)

	for k, v := range nums {
		if nums[k] > 0 {
			if !dedup[v] {
				dedup[v] = true

				if min > v {
					min = v
				}

				if max < v {
					max = v
				}

				xs = append(xs, nums[k])
			}
		}
	}

	n = len(xs)

	if n == 0 {
		return 1
	}


	if min > 1 {
		return 1
	}

	if n == max {
		return max + 1
	}

	// we have to find a hole
	// search the possible range [1,.., n]
	// by splitting range and counting how many fall in each half
	// then splitting that half etc

	return hole(xs, 1, n)
}
