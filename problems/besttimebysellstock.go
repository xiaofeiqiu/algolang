package main

import "math"

func maxProfit(prices []int) int {
	if prices == nil {
		return 0
	}

	minP := math.MaxInt
	res := 0

	for _, v := range prices {
		if v < minP {
			minP = v
			continue
		}

		if v-minP > res {
			res = v - minP
		}
	}

	return res
}
