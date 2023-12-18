package main

func twoSum1(nums []int, target int) []int {

	if nums == nil {
		return []int{}
	}

	lookingFor := make(map[int]int)
	// loop over the array, check if num in lookfor, if yes, return indexes
	// else add looking for val to the map
	for i, v := range nums {
		if index, found := lookingFor[v]; found {
			return []int{i, index}
		}

		lookingFor[target-v] = i
	}

	return []int{}
}
