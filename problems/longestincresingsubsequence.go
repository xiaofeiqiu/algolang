package main

func lengthOfLIS(nums []int) int {

	if nums == nil {
		return 0
	}

	n := len(nums)
	dp := make([]int, n)

	dp[0] = 1

	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := -1
	for i := range dp {
		res = max(res, dp[i])
	}
	return res
}

// solution 2

func lengthOfLIS2(nums []int) int {
	tails := []int{}

	for _, num := range nums {
		pos := findJustGreater(tails, num)
		if pos == len(tails) {
			tails = append(tails, num)
		} else {
			tails[pos] = num
		}
	}
	return len(tails)
}

func findJustGreater(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid
		}
	}

	if nums[left] >= target {
		return left
	}

	if nums[right] >= target {
		return right
	}

	return len(nums)
}
