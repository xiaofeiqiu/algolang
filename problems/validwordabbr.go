package main

import (
	"strconv"
	"unicode"
)

func validWordAbbreviation(word string, abbr string) bool {
	wordIndex, abbrIndex := 0, 0

	for abbrIndex < len(abbr) && wordIndex < len(word) {

		if unicode.IsLetter(rune(abbr[abbrIndex])) {
			if abbr[abbrIndex] != word[wordIndex] {
				return false
			}
			wordIndex++
			abbrIndex++
			continue
		}

		if unicode.IsDigit(rune(abbr[abbrIndex])) {
			if abbr[abbrIndex] == '0' {
				return false
			}
			end := abbrIndex
			for end < len(abbr) && unicode.IsDigit(rune(abbr[end])) {
				end++
			}
			num, _ := strconv.Atoi(abbr[abbrIndex:end])
			abbrIndex = end
			wordIndex += num
			continue
		}

		return false
	}

	return wordIndex == len(word) && abbrIndex == len(abbr)
}
