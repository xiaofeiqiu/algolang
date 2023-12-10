package main

func twoSum(nums []int, target int) []int {

	res := make([]int, 2)
	if nums == nil {
		return res
	}

	lookingFor := make(map[int]int)
	// loop over the array, check if num in lookfor, if yes, return indexes
	// else add looking for val to the map
	for i, v := range nums {
		if index, found := lookingFor[v]; found {
			res[0] = i
			res[1] = index
			return res
		}

		lookingFor[target-v] = i
	}

	return res
}
