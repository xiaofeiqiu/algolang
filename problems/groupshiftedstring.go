package main

import "fmt"

func groupStrings(strs []string) [][]string {
	kv := make(map[string][]string)

	for i := range strs {
		diffs := make([]int, 0)
		for j := 1; j < len(strs[i]); j++ {
			// Calculate the cyclic difference and convert it to a positive value
			diff := (int(strs[i][j]) - int(strs[i][j-1]) + 26) % 26
			diffs = append(diffs, diff)
		}

		key := fmt.Sprintf("%v", diffs)
		kv[key] = append(kv[key], strs[i])
	}

	res := make([][]string, 0)
	for k := range kv {
		res = append(res, kv[k])
	}

	return res
}
