package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	if strs == nil {
		return [][]string{}
	}

	anagramMap := make(map[string][]string)
	for i := range strs {
		// sort the chars
		chars := []rune(strs[i])
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		// use sorted string as ket, append strs[i] to the value
		key := string(chars)
		anagramMap[key] = append(anagramMap[key], strs[i])
	}

	// convert to [][]string
	res := make([][]string, 0)
	for k := range anagramMap {
		res = append(res, anagramMap[k])
	}

	return res
}
