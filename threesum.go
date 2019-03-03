package leet

import "sort"

func ThreeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	res := make([][]int, 0)
	xs := nums
	sort.Sort(sort.IntSlice(xs))
	pi := -1
	for i, v := range xs {
		if xs[i] > 0 {
			break
		}

		if pi > -1 && xs[i] == xs[pi] {
			continue
		}
		pi = i
		pj := -1

		for j := i + 1; j < len(xs) - 1; j++ {
			if v + xs[j] > 0 {
				break
			}
			if pj > -1 && xs[j] == xs[pj] {
				continue
			}
			pj = j
			x := -(v + xs[j])
			k := j + 1 + sort.SearchInts(xs[j + 1:], x)
			if k < len(xs) && xs[k] == x {
				if len(res) > 0 {
					ps := res[len(res) -1]
					if ps[0] != v || ps[1] != xs[j] {
						res = append(res, []int{v, xs[j], xs[k]})
					}
				} else {
					res = append(res, []int{v, xs[j], xs[k]})
				}
			}
		}
	}

	return res
}

