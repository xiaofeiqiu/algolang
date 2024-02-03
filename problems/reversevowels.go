package main

import "strings"

func reverseVowels(s string) string {

	i := 0
	j := len(s) - 1

	bytes := []byte(s)

	for i < j {
		if !isVowel(bytes[i]) {
			i++
			continue
		}

		if !isVowel(bytes[j]) {
			j--
			continue
		}

		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	return string(bytes)
}

func isVowel(b byte) bool {
	return strings.Contains("aeiouAEIOU", string(b))
}
