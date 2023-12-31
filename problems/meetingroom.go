package main

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	if intervals == nil {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i, v := range intervals {
		if i == 0 {
			continue
		}

		if intervals[i-1][1] > v[0] {
			return false
		}
	}

	return true
}
