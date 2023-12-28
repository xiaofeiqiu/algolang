package main

type NestedInteger interface {
	IsInteger() bool
	GetInteger() int
	GetList() []NestedInteger
}

func depthSum(nestedList []NestedInteger) int {
	return dfs2(nestedList, 1)
}

func dfs2(nestedList []NestedInteger, depth int) int {
	if len(nestedList) == 0 {
		return 0
	}

	sum := 0
	for i := range nestedList {
		if nestedList[i].IsInteger() {
			sum += nestedList[i].GetInteger() * depth
		} else {
			sum += dfs2(nestedList[i].GetList(), depth+1)
		}
	}

	return sum
}
