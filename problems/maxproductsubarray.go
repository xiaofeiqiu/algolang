package main

func maxProduct(nums []int) int {

	if nums == nil {
		return 0
	}

	res := nums[0]
	currMin, currMax := nums[0], nums[0]
	for i := range nums {
		if i == 0 {
			continue
		}

		if nums[i] < 0 {
			currMin, currMax = currMax, currMin
		}
		currMin = min(currMin*nums[i], nums[i])
		currMax = max(currMax*nums[i], nums[i])
		res = max(res, currMax)
	}

	return res
}
