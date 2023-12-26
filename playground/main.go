package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	l := 2
	h := l
	k := 3
	sum := 0

	count := 0
	for i := range nums {
		if h > 0 {
			sum += nums[i]
			h--
			continue
		}

		if sum == k {
			count++
		}

		sum += (nums[i] - nums[i-l])
		fmt.Printf("i=%d, i-l=%d\n", nums[i], nums[i-l])
	}

	fmt.Println(count)
}
