package main

func maxArea(height []int) int {
	if height == nil {
		return 0
	}

	left, right := 0, len(height)-1
	res := -1
	for left < right {
		h := min(height[left], height[right])
		res = max(res, h*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}
