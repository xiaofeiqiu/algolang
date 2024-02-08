package main

import "sort"

func maximumPopulation(logs [][]int) int {

	yearPopulation := map[int]int{}

	for _, v := range logs {
		yearPopulation[v[0]]++
		yearPopulation[v[1]]--
	}

	years := make([]int, 0)
	for year := range yearPopulation {
		years = append(years, year)
	}

	sort.Ints(years)

	res, count, maxCount := 0, 0, 0
	for _, v := range years {
		count += yearPopulation[v]
		if count > maxCount {
			maxCount = count
			res = v
		}
	}

	return res
}
