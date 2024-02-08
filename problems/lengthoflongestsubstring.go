package main

func lengthOfLongestSubstring(s string) int {

	if s == "" {
		return 0
	}

	getIndex := make(map[byte]int, 0)

	left := 0
	res := 0

	for right := range s {
		// condition to shrink the window: if repeated found
		// and the repeated index is within the window
		// shrink left pointer to exclude the char
		if index, found := getIndex[s[right]]; found && index >= left {
			left = index + 1
		}

		// update the chars location
		getIndex[s[right]] = right

		// cal result
		res = max(res, right-left+1)
	}

	return res
}
