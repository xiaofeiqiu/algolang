package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	l, r := 0, len(s)-1
	s = strings.ToLower(s)
	chars := []rune(s)
	for l < r {

		if !unicode.IsLetter(chars[l]) && !unicode.IsNumber(chars[l]) {
			l++
			continue
		}

		if !unicode.IsLetter(chars[r]) && !unicode.IsNumber(chars[r]) {
			r--
			continue
		}

		if chars[l] == chars[r] {
			l++
			r--
		} else {
			return false
		}
	}

	return true
}
