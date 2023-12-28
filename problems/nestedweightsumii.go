package main

// solution 1
func depthSumInverse(nestedList []NestedInteger) int {
	maxDepth := getDepth(nestedList)
	return dfsnested(nestedList, maxDepth)
}

func getDepth(nestedList []NestedInteger) int {
	if len(nestedList) == 0 {
		return 0
	}

	depth := 0
	// compute depth for each child and get maxdepth amoung all childs
	for i := range nestedList {
		depth = max(depth, getDepth(nestedList[i].GetList()))
	}

	// return max child depth + 1
	return depth + 1
}

func dfsnested(nestedList []NestedInteger, weight int) int {
	if len(nestedList) == 0 {
		return 0
	}

	sum := 0
	for i := range nestedList {
		if nestedList[i].IsInteger() {
			sum += nestedList[i].GetInteger() * weight
		} else {
			sum += dfsnested(nestedList[i].GetList(), weight-1)
		}
	}

	return sum
}

// solution 2

func depthSumInverse1(nestedList []NestedInteger) int {

	if nestedList == nil {
		return 0
	}

	runningSum := 0
	result := 0

	queue := append([]NestedInteger{}, nestedList...)

	for len(queue) > 0 {
		size := len(queue)

		levelSum := 0
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr.IsInteger() {
				levelSum += curr.GetInteger()
				continue
			}

			if len(curr.GetList()) > 0 {
				queue = append(queue, curr.GetList()...)
			}
		}

		runningSum += levelSum
		result += runningSum
	}

	return result
}
