package main

import "math"

func maxPoints(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}

	max := 0

	// for each point, calculate slop with every other point
	// number of points in the same line = slop occurence + overlap points + 1 (itself)
	for i := 0; i < len(points); i++ {
		// maintain a map to count slop occurence
		slopCount := make(map[float64]int)
		overlap := 0
		for j := i + 1; j < len(points); j++ {

			// if point overlap
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				overlap++
				continue
			}

			slop := calSlop(points[i], points[j])
			slopCount[slop]++

			// if local max points > max, update result
			if slopCount[slop]+overlap+1 > max {
				max = slopCount[slop] + overlap + 1
			}
		}
	}

	return max
}

func calSlop(p1, p2 []int) float64 {
	if p1[0] == p2[0] {
		return math.Inf(1)
	}
	deltaX := p1[0] - p2[0]
	deltaY := p1[1] - p2[1]
	return float64(deltaY) / float64(deltaX)
}
