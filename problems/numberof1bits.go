package main

func hammingWeight(num uint32) int {
	res := 0

	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			res++
		}
		num = num >> 1
	}

	return res
}
