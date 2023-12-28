package main

// solution 1
func findBuildings(heights []int) []int {

	max := -1
	res := make([]int, 0)
	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > max {
			max = heights[i]
			res = append(res, i)
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

// solution 2
func findBuildings2(heights []int) []int {

	nextGreater := make([]int, len(heights))
	stack := make([]int, 0)
	for i := range nextGreater {
		nextGreater[i] = -1
	}

	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[i] > heights[top(stack)] {
			pop(&stack)
		}
		if len(stack) > 0 {
			nextGreater[i] = top(stack)
		}
		stack = append(stack, i)
	}

	res := make([]int, 0)
	for i, v := range nextGreater {
		if v == -1 {
			res = append(res, i)
		}
	}
	return res
}

func top(stack []int) int {
	return stack[len(stack)-1]
}

func pop(stack *[]int) int {
	num := top(*stack)
	*stack = (*stack)[:len(*stack)-1]
	return num
}
