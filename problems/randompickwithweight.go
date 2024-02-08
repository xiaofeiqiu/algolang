package main

import "math/rand"

type Solution struct {
	prefix []int
}

func Constructor12(w []int) Solution {
	prefix := make([]int, 0)
	sum := 0
	for i := range w {
		sum += w[i]
		prefix = append(prefix, sum)
	}

	return Solution{prefix}
}

func (this *Solution) PickIndex() int {
	target := rand.Intn(this.prefix[len(this.prefix)-1]) + 1

	left, right := 0, len(this.prefix)-1
	for left <= right {
		mid := left + (right-left)/2
		if this.prefix[mid] == target {
			return mid
		} else if this.prefix[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
