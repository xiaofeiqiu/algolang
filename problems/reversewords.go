package main

import "strings"

func reverseWords(s string) string {

	if s == "" {
		return ""
	}

	s = strings.Trim(s, " ")
	words := strings.Fields(s) // extract word array

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
