package main

import (
	"strings"
	"unicode"
)

type Encoder struct {
	Str   string
	Count int
}

func decodeString(s string) string {

	chars := []rune(s)
	count := 0
	str := ""
	stack := make([]Encoder, 0)

	for _, v := range chars {
		if unicode.IsLetter(v) {
			str += string(v)
			continue
		}

		if unicode.IsDigit(v) {
			count = count*10 + int(v-'0')
			continue
		}

		if v == '[' {
			stack = append(stack, Encoder{str, count})
			str = ""
			count = 0
			continue
		}

		if v == ']' {
			e := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			str = e.Str + strings.Repeat(str, e.Count)
		}
	}

	return str
}
