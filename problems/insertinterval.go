package main

const end = 1
const start = 0

func insert(intervals [][]int, newInterval []int) [][]int {
	if intervals == nil || newInterval == nil {
		return intervals
	}

	res := make([][]int, 0)

	i := 0
	// copy non overlap intervals
	// if current end time < newStartTime, then no overlap
	for ; i < len(intervals) && intervals[i][end] < newInterval[start]; i++ {
		res = append(res, intervals[i])
	}

	// start merging process
	// merge if overlap exist: current interval start time <= new interval end time
	// create a merged var that copies the new interval's value
	merged := []int{newInterval[start], newInterval[end]}
	for ; i < len(intervals) && intervals[i][start] <= newInterval[end]; i++ {
		merged[start] = min(intervals[i][start], merged[start])
		merged[end] = max(intervals[i][end], merged[end])
	}

	// append merged interval
	res = append(res, merged)

	// copy the rest of the intervals
	for ; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}

	return res
}
