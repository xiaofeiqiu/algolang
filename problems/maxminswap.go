package main

import "strconv"

func maximumSwap(num int) int {
	chs := []rune(strconv.Itoa(num))
	n := len(chs)

	maxIndex := n - 1
	switchFrom, switchTo := -1, -1

	for i := n - 1; i >= 0; i-- {
		if chs[i] > chs[maxIndex] {
			maxIndex = i
		}

		if chs[maxIndex] > chs[i] {
			switchFrom = i
			switchTo = maxIndex
		}
	}

	if switchFrom != -1 {
		chs[switchFrom], chs[switchTo] = chs[switchTo], chs[switchFrom]
	}

	res, _ := strconv.Atoi(string(chs))
	return res
}
