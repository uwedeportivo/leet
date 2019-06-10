package leet

import "sort"

func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	startIndex := sort.Search(len(intervals), func(i int) bool {
		return intervals[i][0] >= newInterval[0]
	})

	endIndex := sort.Search(len(intervals), func(i int) bool {
		return intervals[i][0] > newInterval[1]
	})

	var res [][]int

	for i := 0; i < startIndex - 1; i++ {
		res = append(res, intervals[i])
	}

	adjustedNewInterval := make([]int, 2)
	adjustedNewInterval[0] = newInterval[0]
	adjustedNewInterval[1] = newInterval[1]

	if startIndex > 0 {
		if newInterval[0] <= intervals[startIndex-1][1] {
			adjustedNewInterval[0] = intervals[startIndex-1][0]
		} else {
			res = append(res, intervals[startIndex-1])
		}
	}

	if endIndex > 0 && newInterval[1] < intervals[endIndex-1][1] {
		adjustedNewInterval[1] = intervals[endIndex-1][1]
	}

	res = append(res, adjustedNewInterval)

	for i := endIndex; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}

	return res
}
