package main

import "sort"

const EndTime = 1
const StartTime = 0

func mergeInterval(intervals [][]int) [][]int {
	if intervals == nil {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][StartTime] < intervals[j][StartTime]
	})

	res := make([][]int, 0)
	for i, v := range intervals {
		if i == 0 {
			res = append(res, v)
			continue
		}

		if res[len(res)-1][EndTime] < v[StartTime] {
			res = append(res, v)
		} else {
			res[len(res)-1][EndTime] = max(res[len(res)-1][EndTime], v[EndTime])
		}
	}

	return res
}
