package leet

import "sort"

type intWithIndex struct {
	val int
	index int
}

type byVal []intWithIndex

func (a byVal) Len() int           { return len(a) }
func (a byVal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byVal) Less(i, j int) bool { return a[i].val < a[j].val }


func TwoSum(nums []int, target int) []int {
	xs := make(byVal, len(nums))

	for i,v := range nums {
		xs[i].val = v
		xs[i].index = i
	}

	sort.Sort(xs)

	for k,v := range nums {
		w := target - v
		j := sort.Search(len(xs), func(i int) bool { return xs[i].val >= w })
		if j < len(xs) && xs[j].val == w && xs[j].index != k {
			if k < xs[j].index {
				return []int{k, xs[j].index}
			} else {
				return []int{xs[j].index, k}
			}
		}
	}
	return nil
}

