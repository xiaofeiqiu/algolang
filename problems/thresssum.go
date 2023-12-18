package main

import "sort"

func threeSum(nums []int) [][]int {
	if nums == nil {
		return make([][]int, 0)
	}

	sort.Ints(nums)

	n := len(nums)
	res := make([][]int, 0)
	for i, v := range nums {
		if i == 0 || nums[i] != nums[i-1] {
			comb := twoSum(nums, v, i+1, n-1)
			if len(comb) > 0 {
				res = append(res, comb...)
			}
		}
	}

	return res
}

func twoSum(nums []int, num1 int, l, r int) [][]int {

	res := make([][]int, 0)

	for l < r {
		sum := num1 + nums[l] + nums[r]

		if sum == 0 {
			res = append(res, []int{num1, nums[l], nums[r]})
			l++
			r--

			for l < r && nums[l] == nums[l-1] {
				l++
			}

			for l < r && nums[r] == nums[r+1] {
				r--
			}
			continue
		}

		if sum > 0 {
			r--
		} else {
			l++
		}
	}

	return res
}
