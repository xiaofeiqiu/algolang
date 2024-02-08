package main

type SparseVector struct {
	nonZeroMap map[int]int
}

func Constructor11(nums []int) SparseVector {
	kv := map[int]int{}
	for i, v := range nums {
		if v != 0 {
			kv[i] = v
		}
	}
	return SparseVector{kv}
}

// Return the dotProduct of two sparse vectors
func (s *SparseVector) dotProduct(vec2 SparseVector) int {

	shorter := s.nonZeroMap
	longer := vec2.nonZeroMap

	if len(shorter) > len(longer) {
		shorter = vec2.nonZeroMap
		longer = s.nonZeroMap
	}

	sum := 0
	for k, v1 := range shorter {
		if v2, exist := longer[k]; exist {
			sum += v1 * v2
		}
	}

	return sum
}
