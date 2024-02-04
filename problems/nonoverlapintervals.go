package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {

	if intervals == nil {
		return 0
	}

	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	end := intervals[0][1]

	count := 0
	for i, v := range intervals {
		if i == 0 {
			continue
		}
		if v[0] < end {
			count++
		} else {
			end = v[1]
		}
	}

	return count
}
