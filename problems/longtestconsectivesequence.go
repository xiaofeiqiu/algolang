package main

func longestConsecutive(nums []int) int {
	if nums == nil {
		return -1
	}

	set := make(map[int]bool)
	for i := range nums {
		set[nums[i]] = true
	}

	res := 0
	// iterate through the set to prevent duplicate processing
	for num := range set {
		// check num - 1 exist or not, it not, this is the starting of the sequence
		if !set[num-1] {
			// keep checking next number exist or not
			l := 1
			next := num + 1
			for set[next] {
				next++
				l++
			}
			// update max length
			if l > res {
				res = l
			}
		}
	}

	return res
}
