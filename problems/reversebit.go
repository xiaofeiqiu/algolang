package main

func reverseBits(num uint32) uint32 {
	if num == 0 {
		return 0
	}

	var res uint32 = 0
	// res should shift left before "or"
	// this is because if we "or" before shift the last element will be shifted 1 distance to the left
	// but we do not need to shift the last bit
	for i := 0; i < 32; i++ {
		res = (res << 1) | (num & 1)
		num = num >> 1
	}

	return res
}
